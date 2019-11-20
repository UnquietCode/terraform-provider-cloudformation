package plugin

import (
	"errors"
	"os"
  "log"
	"fmt"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "crypto/md5"
)



func readResource(id string, meta ProviderMetadata) (map[string]interface{}, string, error) {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
	
  var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	exists, err := fileExists(path)
	
	if err != nil {
		return nil, EMPTY, err
	}
	
	if !exists {
		return nil, EMPTY, nil
	}
	
	// try to read the file
	rawData, hashcode, err := readAndHashFile(path)
	
	if err != nil {
		return nil, EMPTY, err
	}
	
	// turn the data into schema data
	var data map[string]interface{}

  if err := json.Unmarshal(rawData, &data); err != nil {
    return nil, EMPTY, err
  }
	
	return data, hashcode, nil
}


func writeResource(logicalId string, data map[string]interface{}, meta ProviderMetadata) (string, error) {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
    
  // write the resource file
  var path string = fmt.Sprintf("%s/%s.json", meta.workdir, logicalId)  
  hash, err := writeFile(path, data)
  
  if err != nil {
    return EMPTY, err
  }
  
  // update the index file
  readAndWriteFile(TEMPLATE_INDEX_FILE, meta, func(indexData map[string]interface{}) {
    if _, ok := indexData["resources"]; !ok {
      indexData["resources"] = map[string]interface{}{}
    }
    indexData["resources"].(map[string]interface{})[logicalId] = hash
  })
  
  return hash, nil
}

func removeResource(logicalId string, meta ProviderMetadata) error {
	meta.mutex.Lock(LOCK_RESOURCE_READ_WRITE)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_WRITE)
    
  // remove the resource file
  var err error = removeFile(logicalId, meta)
  
  if err != nil {
    return err
  }
  
  // update the index file
  _, _, err = readAndWriteFile(TEMPLATE_INDEX_FILE, meta, func(data map[string]interface{}) {
    if _, ok := data["resources"]; ok {
      resources := data["resources"].(map[string]interface{})
      
      if _, ok := resources["logicalId"]; ok {
        delete(resources, "logicalId")
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
    if err := json.Unmarshal(rawData, &data); err != nil {
      return EMPTY, EMPTY, err
    }
  }
  
  // alter the file
  modifier(data)
  
  // write the results
  newHash, err := writeFile(path, data)
  
	return oldHash, newHash, nil
}


func markResourceAsRead(logicalId string, changeType ChangeType, replaceIndex bool, meta ProviderMetadata) error {
  log.Printf("------------------ read ------- %s", logicalId)
  
	meta.mutex.Lock(LOCK_RESOURCE_READ_COUNT)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_COUNT)

  var err error
  
  if replaceIndex {
    if (*meta.newIndex)[TEMPLATE_COUNTER_FILE] == true {
      removeFile(TEMPLATE_COUNTER_FILE, meta)
      (*meta.newIndex)[TEMPLATE_COUNTER_FILE] = false
    }
  }
  
  _, _, err = readAndWriteFile(TEMPLATE_COUNTER_FILE, meta, func(indexData map[string]interface{}) {
    ensureArrays(indexData, Changed, Unchanged, Maybe)
    
    var addTo string = changeType
    // var removes = allWithout(changeType, Changed, Unchanged, Maybe)
    
    indexData[addTo] = addString(logicalId, indexData[addTo].([]interface{}))
    // indexData[removes[0]] = removeString(logicalId, indexData[removes[0]].([]interface{}))
    // indexData[removes[1]] = removeString(logicalId, indexData[removes[1]].([]interface{}))
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
	
	data, hash, err := readResource(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	// doesn't exist
	if data == nil && hash == EMPTY {
		resourceData.SetId("")
		return nil
	}
    
  // file has changed
  if resourceData.Id() != hash {
  	resourceData.SetId("")
    return nil
  }

  log.Printf("------------------ reading %s as id ------- %s", logicalId, resourceData.Id())
  // log.Printf("------------------ %s", resourceData)
  
  mergeMapToResource("", resourceData, data)
  markResourceAsRead(logicalId, Maybe, true, providerMeta)
  
  return nil
}


func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	
	var data map[string]interface{} = convertResourceToMap("", resourceSchema, resourceData, normalGetter)
	hash, err := writeResource(logicalId, data, providerMeta)
	
	if err != nil {
		return err
	}
  
	resourceData.SetId(hash)
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
	
	resourceData.SetId("")	
  return nil
}


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
  log.Printf("--------------------- diffing, %s %s", logicalId, changes)
  
  var changeType ChangeType
  
  if len(changes) > 0 {
    changeType = Changed
  } else {
    changeType = Unchanged
  }
  markResourceAsRead(logicalId, changeType, false, providerMeta)
	
	return nil
}