package plugin

import (
	"os"
  "log"
  "errors"
  "time"
	"fmt"
	"sort"
	"encoding/json"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"  
)


func buildTemplateReference(meta ProviderMetadata) (map[string]TemplateEntry, string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, TEMPLATE_INDEX_FILE)
	exists, err := fileExists(path)
  
  if err != nil {
    return nil, EMPTY, err
  }
  
  if !exists {
    return map[string]TemplateEntry{}, EMPTY, nil
  }
    
  file, err := os.Open(path)
  
	if err != nil {
      return nil, EMPTY, err
  }
  defer file.Close()

  // read the file
	rawData, _, err := readAndHashFile(path)
	
	if err != nil {
    return nil, EMPTY, err
  }
  
  // convert to data
	var data map[string]interface{} = map[string]interface{}{}

  if err := json.Unmarshal(rawData, &data); err != nil {
    return nil, EMPTY, err
  }  
  
  // build sorted list of keys
  var resources map[string]TemplateEntry = map[string]TemplateEntry{}
  
  if _, ok := data["resources"]; ok {
    for k,v := range data["resources"].(map[string]interface{}) {
      resources[k] = TemplateEntry{
        CfnType: v.(map[string]interface{})["type"].(string),
        Hash: v.(map[string]interface{})["hash"].(string),
      }
    }    
  }
  
	var keys []string
  
	for key, _ := range resources {
    keys = append(keys, key)
	}
	
	sort.Strings(keys)
	
	var bigHash string = ""
	
	for _, key := range keys {
		bigHash += resources[key].Hash
	}
	
	bigHash = strconv.Itoa(hashcode.String(bigHash))
	return resources, bigHash, nil
}


func readCounts(meta ProviderMetadata) (map[ChangeType][]string, error) {
	meta.mutex.Lock(LOCK_RESOURCE_READ_COUNT)
	defer meta.mutex.Unlock(LOCK_RESOURCE_READ_COUNT)
  
  // check for existence
  var path string = fmt.Sprintf("%s/%s.json", meta.workdir, TEMPLATE_COUNTER_FILE)
  exists, err := fileExists(path)
  
  if err != nil {
    return nil, err
  }
  
  if !exists {
		return nil, nil
  }
    
	// turn the data into schema data
  var data map[string]interface{} = map[string]interface{}{}
	
	// try to read the file
	rawData, _, err := readAndHashFile(path)
  
	if err != nil {
    return nil, err
	}
  
  if err := json.Unmarshal(rawData, &data); err != nil {
    return nil, err
  }

  results := map[ChangeType][]string{}
  ensureArraysCT(results, Changed, Unchanged, Maybe)
  
  for _, changeType := range []ChangeType{Changed, Unchanged, Maybe} {
    for _, value := range data[changeType].([]interface{}) {
      results[changeType] = append(results[changeType], value.(string))
    }
  }
  
  return results, nil
}


func extractDeletes(indexData map[ChangeType][]string) []string {
  for _, maybe := range indexData[Maybe] {

    // maybe - changed - unchanged = deletes
    if !arrayContainsString(indexData[Changed], maybe) && !arrayContainsString(indexData[Unchanged], maybe) {
      indexData[Maybe] = removeFromStringArray(maybe, indexData[Maybe])
      indexData[Deleted] = append(indexData[Deleted], maybe)
    }
  }
  
  return indexData[Deleted]
}


func waitForChangesNew(meta ProviderMetadata) (bool, error) {  
  createStateConf := &resource.StateChangeConf{
    Pending: []string{
      STATE_WAIT,
    },
    Target: []string{
      STATE_DONE,
    },
    Refresh: func() (interface{}, string, error) {
      refs, _, err := buildTemplateReference(meta)
      if err != nil { return nil, EMPTY, err }
      
      countLists, err := readCounts(meta)
      if err != nil { return nil, EMPTY, err }
  
      // wait for list to create
      if countLists == nil {
        return nil, STATE_WAIT, nil
      }
      
      refCount := len(refs)
      deletes := extractDeletes(countLists) // has to go first!
      changed := countLists[Changed]
      unchanged := countLists[Unchanged]
      maybeChanged := countLists[Maybe]
      
      counts := len(changed) + len(unchanged) + len(deletes) + len(maybeChanged)
      log.Printf("waiting for template refs to converge on create (%d / %d) %s", refCount, counts, countLists)

      if counts == refCount {
        return (len(changed) + len(deletes)) > 0, STATE_DONE, nil
      }
      
      return nil, STATE_WAIT, nil
    },
    Timeout: 5 * time.Minute,
    // Delay: 1 * time.Second,
  }

  isChanged, err := createStateConf.WaitForState()
  
  if err != nil {
    return false, err
  }
  
  return isChanged.(bool), nil
}


func waitForChanges(deleterable bool, meta ProviderMetadata) (bool, error) {  
  refs, refHash, err := buildTemplateReference(meta)
  
  if err != nil {
    return false, err
  }
  
  refCount := len(refs)
  
  if refCount == 0  &&  refHash == EMPTY {
    return false, nil
  }

  createStateConf := &resource.StateChangeConf{
    Pending: []string{
      STATE_WAIT,
    },
    Target: []string{
      STATE_DONE,
    },
    Refresh: func() (interface{}, string, error) {    
      countLists, err := readCounts(meta)
      if err != nil { return nil, EMPTY, err }
      
      // wait for list to create
      if countLists == nil {
        return nil, STATE_WAIT, nil
      }
      
      var counts int = 0
      var changes int = 0
      
      if deleterable {
        deletes := extractDeletes(countLists) // has to go first!
        changed := countLists[Changed]
        unchanged := countLists[Unchanged]

        counts += len(changed) + len(unchanged) + len(deletes)
        changes += len(changed) + len(deletes)
      } else {
        maybeChanged := countLists[Maybe]
        changed := countLists[Changed]
        unchanged := countLists[Unchanged]        
        
        counts += len(maybeChanged) + len(changed) + len(unchanged)
        changes += len(changed)
      }
      
      log.Printf("waiting for template refs to converge (%d / %d) %s", refCount, counts, countLists)

      if changes > 0 {
        return true, STATE_DONE, nil
      }      
      
      if counts == refCount {
        return changes > 0, STATE_DONE, nil
      }
      
      if counts >= refCount {
        return nil, EMPTY, errors.New("too many refs?")
      }

      return nil, STATE_WAIT, nil
    },
    Timeout: 5 * time.Minute,
    // Delay: 1 * time.Second,
  }

  isChanged, err := createStateConf.WaitForState()
  
  if err != nil {
    return false, err
  }
  
  return isChanged.(bool), nil
}


func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)    
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, RENDERED_TEMPLATE_FILE)
	
  // check if it exists
  exists, err := fileExists(path)  
	if err != nil { return err }
  
	if !exists {
		resourceData.SetId("")
    return nil
	}
  
  // check if the file hash is the same
	data, hash, err := readAndHashFile(path)
  resourceData.Set("output", string(data))

  if resourceData.Id() != hash {
		resourceData.SetId("")
    return nil
  }
  
  _, err = waitForChanges(false, meta.(ProviderMetadata))
	if err != nil { return err }
  // 
  // if isChanged {
  //   // resourceDiff.SetNewComputed("hash")
  //   resourceData.Set("hash", "0")
  // }  
  
  return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)  
	
  _, err := waitForChangesNew(meta.(ProviderMetadata))
	if err != nil { return err }
  
  resourceHashes, _, err := buildTemplateReference(providerMeta)
	
  if err != nil {
		return err
	}
  
	var templateData map[string]interface{} = map[string]interface{}{}
	var resources = map[string]interface{}{}
	templateData["Resources"] = resources
  
	for name, entry := range resourceHashes {
		properties, _, err := readResource(name, providerMeta)
  
		if err != nil {
			return err
		}
    
    if properties == nil {
      continue
    }
  
		data := map[string]interface{}{}
		data["Type"] = entry.CfnType
  
		resources[name] = data
    delete(properties, "logical_id")
  
    data["Properties"] = deSnake(properties)
	}
  
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, RENDERED_TEMPLATE_FILE)
	hash, content, err := writeFilePlusContent(path, templateData)
  
	if err != nil {
		return err
	}
  
	resourceData.SetId(hash)  
  resourceData.Set("output", content)
	
  return nil
}


func TemplateUpdate(resourceData *schema.ResourceData, meta interface{}) error {
  return TemplateCreate(resourceData, meta)
}


func TemplateDelete(resourceData *schema.ResourceData, meta interface{}) error {
	err := removeFile(RENDERED_TEMPLATE_FILE, meta.(ProviderMetadata))
  
	if err != nil {
		return err
	}
  
	resourceData.SetId("")
  return nil
}

func TemplateCustomizeDiff(resourceDiff *schema.ResourceDiff, meta interface{}) error {
  isChanged, err := waitForChanges(true, meta.(ProviderMetadata))
	if err != nil { return err }
  
  if isChanged {
    resourceDiff.SetNewComputed("hash")
  }  
  
	return nil
}
