package plugin

import (
	"errors"
	"os"
  "log"
	"fmt"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


func countResources(meta ProviderMetadata) (int, error) {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
  
  var path string = fmt.Sprintf("%s/%s.json", meta.workdir, RENDERED_TEMPLATE_FILE)

  exists, err := fileExists(path)
  if err != nil { return -1, err }
  
  if !exists {
    return -1, nil
  }

	rawData, _, err := readAndHashFile(path)
	if err != nil { return -1, err }

  if string(rawData) == "" {
    return -1, nil
  }
  
  var data map[string]interface{} = map[string]interface{}{}

  if err := json.Unmarshal(rawData, &data); err != nil {
    return -1, err
  }
    
  if _, ok := data["Resources"]; !ok {
    return -1, nil
  }
  
  count := len(data["Resources"].(map[string]interface{}))
  return count, nil
}

func writeResource(resourceType string, logicalId string, data map[string]interface{}, meta ProviderMetadata) (string, error) {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
  
  readAndWriteFile(RENDERED_TEMPLATE_FILE, meta, func(indexData map[string]interface{}) {
    if _, ok := indexData["Resources"]; !ok {
      indexData["Resources"] = map[string]interface{}{}
    }
    
    resourceData := map[string]interface{}{}
		resourceData["Type"] = resourceType
  
    properties := deSnake(data)
    delete(properties, "logical_id")
    resourceData["Properties"] = properties
    
    indexData["Resources"].(map[string]interface{})[logicalId] = resourceData
  })
  
  return "", nil
}

func removeResource(logicalId string, meta ProviderMetadata) error {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
  
  _, _, err := readAndWriteFile(RENDERED_TEMPLATE_FILE, meta, func(data map[string]interface{}) {
    if _, ok := data["Resources"]; ok {
      resources := data["Resources"].(map[string]interface{})
      
      if _, ok := resources[logicalId]; ok {
        delete(resources, logicalId)
      }
    }
  })
  
  if err != nil {
    return err
  }
  
  return nil
}


type DataModifier func(data map[string]interface{})


func readAndWriteFile(fileName string, meta ProviderMetadata, modifier DataModifier) (string, string, error) {	
  var path string = fmt.Sprintf("%s/%s.json", meta.workdir, fileName)
	var file *os.File = nil
  var created bool = false
  log.Printf("")	
  
  // check for existence
  exists, err := fileExists(path)
  
  if err != nil {
    return EMPTY, EMPTY, err
  }
  
  if !exists {
		f, err := os.Create(path)
		
		if err != nil {
      return EMPTY, EMPTY, err
		}
		
		file = f
    created = true
  }
  
	// open the file if it wasn't created
	if file == nil {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		
		if err != nil {
	    return EMPTY, EMPTY, err
		}
		
		file = f
  }
  
	// turn the data into schema data
  var data map[string]interface{} = map[string]interface{}{}
  var oldHash string
  
  if !created {
	
  	// try to read the file
  	rawData, oldHashh, err := readAndHashFile(path)
  	oldHash = oldHashh
    
  	if err != nil {
      return EMPTY, EMPTY, err
  	}
    
    if string(rawData) != "" {
      if err := json.Unmarshal(rawData, &data); err != nil {
        return EMPTY, EMPTY, err
      }
    }
  }
  
  // alter the file
  modifier(data)
  
  // write the results
  newHash, err := writeFile(path, data)
  
	return oldHash, newHash, nil
}


func markResourceAsRead(logicalId string, changeType ChangeType, replaceIndex bool, meta ProviderMetadata) error {
  meta.mutex.Lock(LOCK_RESOURCE_READ_COUNT)
  defer meta.mutex.Unlock(LOCK_RESOURCE_READ_COUNT)

  var err error
  
  if replaceIndex {
    var path string = fmt.Sprintf("%s/%s.json", meta.workdir, TEMPLATE_COUNTER_FILE)
    
    if (*meta.newIndex)[TEMPLATE_COUNTER_FILE] == true {
      err = waitForEmptyFile(path)
      if err != nil { return err }
      (*meta.newIndex)[TEMPLATE_COUNTER_FILE] = false
    }    
  }
  
  _, _, err = readAndWriteFile(TEMPLATE_COUNTER_FILE, meta, func(indexData map[string]interface{}) {
    ensureArrays(indexData, Unchanged, Changed, Maybe, Deleted, Updated)
    
    var addTo string = changeType
    indexData[addTo] = addString(logicalId, indexData[addTo].([]interface{}))
    
    for _, maybe := range indexData[Maybe].([]interface{}) {
    
      // remove changed from maybe
      if arrayContainsItem(indexData[Changed], maybe) {
        indexData[Maybe] = removeFromChangesArray(maybe, indexData[Maybe])
        continue
      }
    
      // remove unchanged from maybe
      if arrayContainsItem(indexData[Unchanged], maybe) {
        indexData[Maybe] = removeFromChangesArray(maybe, indexData[Maybe])
      }
    }
  })
  
  if err != nil {
    return err
  }
  
  return nil
}



// --------------------------	------------------------------------------------

func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
  
  markResourceAsRead(logicalId, Maybe, true, providerMeta)
  return nil
}


func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	
	var data map[string]interface{} = convertResourceToMap("", resourceSchema, resourceData, normalGetter)
	_, err := writeResource(resourceType, logicalId, data, providerMeta)
	
	if err != nil {
		return err
	}
  
	resourceData.SetId(logicalId)
  markResourceAsRead(logicalId, Updated, false, providerMeta)
  
  return nil
}


func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
  
	return ResourceCreate(resourceType, resourceSchema, resourceData, meta)
}


func ResourceDelete(resourceType string, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	err := removeResource(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
  markResourceAsRead(logicalId, Deleted, false, providerMeta)
	
	resourceData.SetId("")	
  return nil
}
// 
// 
// -func ResourceExists(resourceData *schema.ResourceData, meta interface{}) (bool, error) {
// -	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
// -	var logicalId string = resourceData.Get("logical_id").(string)
// -	exists, err := fileExists(path);	
// -
// -	if err != nil {
// -		return false, err
// -	}
// -	
// -	(*providerMeta.exists)[logicalId] = exists
// -	
// -  return exists, nil
// -}


func ResourceCustomizeDiff(resourceType string, resourceDiff *schema.ResourceDiff, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	logicalId := resourceDiff.Get("logical_id").(string)
	
	if _, ok := (*providerMeta.exists)[logicalId]; ok {
    return errors.New("exists, duplicate logical id "+logicalId)
	}

	if _, ok := (*providerMeta.diffed)[logicalId]; ok {
    return errors.New("diffed, duplicate logical id "+logicalId)
	}
 
  changes := resourceDiff.GetChangedKeysPrefix("")
  // log.Printf("--------------------- diffing, %s %s", logicalId, changes)
  // 
  var changeType ChangeType
  
  if len(changes) > 0 {
    changeType = Changed
  } else {
    changeType = Unchanged
  }
  
  markResourceAsRead(logicalId, changeType, false, providerMeta)
  (*providerMeta.diffed)[logicalId] = true

	return nil
}