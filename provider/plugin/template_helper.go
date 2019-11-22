package plugin

import (
	// "os"
  "log"
  "errors"
  "time"
	"fmt"
	// "sort"
	"encoding/json"
	// "strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"  
)


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
  
  // TODO need to handle no resources
  if string(rawData) == "" {
    return nil, nil
  }
  
  results := map[ChangeType][]string{}
  ensureArraysCT(results, Maybe, Changed, Unchanged, Updated, Deleted)
  
  
  if err := json.Unmarshal(rawData, &data); err != nil {
    return nil, err
  }
  
  for _, changeType := range []ChangeType{Changed, Unchanged, Maybe} {
    for _, value := range data[changeType].([]interface{}) {
      results[changeType] = append(results[changeType], value.(string))
    }
  }
  
  return results, nil
}


func waitForChanges(meta ProviderMetadata) (bool, error) {  
  
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
      
      refs, err := countResources(meta)
      if err != nil { return nil, EMPTY, err }
      
      deletes := countLists[Deleted]
      updated := countLists[Updated]
      changed := countLists[Changed]
      unchanged := countLists[Unchanged]
      maybe := countLists[Maybe]
      
      var counts int = len(maybe) + len(changed) + len(unchanged)
      var changes int = len(updated) + len(deletes)
      
      log.Printf("waiting for template refs to converge (%d / %d) %s", refs, counts, countLists)

      if refs == counts {
        return changes > 0, STATE_DONE, nil
      }
      
      if changes >= counts {
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
  var path string = fmt.Sprintf("%s/%s.json", meta.(ProviderMetadata).workdir, TEMPLATE_COUNTER_FILE)
  writeEmptyFile(path)
  
  return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {
  var providerMeta ProviderMetadata = meta.(ProviderMetadata)
  
  _, err := waitForChanges(providerMeta)
	if err != nil { return err }
  
  var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, RENDERED_TEMPLATE_FILE)
  rawData, _, err := readAndHashFile(path)
  
  resourceData.Set("output", string(rawData))
  resourceData.SetId(RENDERED_TEMPLATE_FILE)
  return nil
}


func TemplateUpdate(resourceData *schema.ResourceData, meta interface{}) error {
  return TemplateCreate(resourceData, meta)
}


func TemplateDelete(resourceData *schema.ResourceData, meta interface{}) error {
	resourceData.SetId("")
  return nil
}


func TemplateCustomizeDiff(resourceDiff *schema.ResourceDiff, meta interface{}) error {
  countLists, err := readCounts(meta.(ProviderMetadata))
  if err != nil { return err }
  
  changed := countLists[Changed]
  maybe := countLists[Maybe]
  isChanged := len(maybe) + len(changed) > 0
  
  if isChanged {
    resourceDiff.SetNewComputed("output")
  }
  
	return nil
}
