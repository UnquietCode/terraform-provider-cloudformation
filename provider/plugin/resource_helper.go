package plugin

import (
	"errors"
	"os"
	"fmt"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


type TemplateEntry struct {
   logicalId string
	 cfnType string
	 hash string
}


func readFile(id string, meta ProviderMetadata) (map[string]interface{}, string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	exists, err := fileExists(path);
	
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


func handleExistingResouce(meta ProviderMetadata, entry TemplateEntry) error {
	meta.mutex.Lock(LOCK_TEMPLATE_REFERENCE)
	defer meta.mutex.Unlock(LOCK_TEMPLATE_REFERENCE)
	
	var file *os.File = nil
	var path string = fmt.Sprintf("%s/template.data.json", meta.workdir)
	
	// if this is the first run, create the file (or overwrite it)
	if (*meta.newIndex == true) {
		f, err := os.Create(path)
		
		if err != nil {
			return err
		}
		
		file = f
		*meta.newIndex = false
	}
	
	// open the file if it wasn't created
	if file == nil {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		
		if err != nil {
			return err
		}
		
		file = f
	}
	
	// append to the file
	var data map[string]interface{} = map[string]interface{}{
	 "id": entry.logicalId,
	 "type": entry.cfnType,
	 "hash": entry.hash,
 	}
	
	entryJson, err := mapToJson(data, false)
	
	if err != nil {
		return err
	}
	
	file.Write(entryJson)
	file.WriteString("\n")
	file.Close()
	
	// append to list
	// *meta.existingResources = append(*meta.existingResources, id)	
	return nil
}

// --------------------------	------------------------------------------------

func ResourceExists(resourceData *schema.ResourceData, meta interface{}) (bool, error) {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, logicalId)

	exists, err := fileExists(path);	

	if err != nil {
		return false, err
	}
	
	(*providerMeta.exists)[logicalId] = exists
	
  return exists, nil
}


func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	
	data, hash, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	// doesn't exist
	if data == nil && hash == EMPTY {
		// handleExistingResouce(providerMeta, logicalId, "")
		resourceData.SetId("")
		return nil
	}
	
	for name, value := range data {
		resourceData.Set(name, value)
	}
	
	resourceData.SetId(hash)
	// handleExistingResouce(providerMeta, logicalId, hash)
	incrementResourceCounter(providerMeta)
	
  return nil
}

func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	
	var data map[string]interface{} = convertResourceToMap("", resourceSchema, resourceData, normalGetter)
	hashcode, err := writeFile(logicalId, data, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId(hashcode)
	incrementResourceCounter(providerMeta)
	
  return nil
}


func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	data, _, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	convertAndMergeResourceToMap("", resourceSchema, resourceData, data, changedOnlyGetter)
	hashcode, err := writeFile(logicalId, data, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId(hashcode)
	incrementResourceCounter(providerMeta)

  return nil
}


func ResourceDelete(resourceType string, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	err := removeFile(logicalId, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId("")
	incrementResourceCounter(providerMeta)
	
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
	
	var entry TemplateEntry = TemplateEntry{
		cfnType: resourceType,
		logicalId: logicalId,
		hash: "---",
	}
	
	handleExistingResouce(providerMeta, entry)
	(*providerMeta.diffed)[logicalId] = true
	
	// log.Printf("customizeResourceDiff %s +++++++++++++++++++++++++++++++++++++++++++++++++", logicalId)
	
	// if id, ok := resourceDiff.GetOk("id"); ok {
	// 	handleExistingResouce(meta.(ProviderMetadata), logical_id.(string), id.(string))
	// } else {
	// 	handleExistingResouce(meta.(ProviderMetadata), logical_id.(string), "")
	// 
	// }
	
	return nil
}