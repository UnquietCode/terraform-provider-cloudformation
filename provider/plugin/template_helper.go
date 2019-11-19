package plugin

import (
	"os"
	"fmt"
	"sort"
	"encoding/json"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)


func buildTemplateReference(meta ProviderMetadata) (map[string]string, string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, TEMPLATE_INDEX_FILE)
	exists, err := fileExists(path)
  
  if err != nil {
    return nil, EMPTY, err
  }
  
  if !exists {
    return map[string]string{}, EMPTY, nil
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
  var resources map[string]string = map[string]string{}
  
  if _, ok := data["resources"]; ok {
    for k,v := range data["resources"].(map[string]interface{}) {
      resources[k] = v.(string)
    }    
  }
  
	var keys []string
  
	for key, _ := range resources {
    keys = append(keys, key)
	}
	
	sort.Strings(keys)
	
	var bigHash string = ""
	
	for _, key := range keys {
		bigHash += resources[key]
	}
	
	bigHash = strconv.Itoa(hashcode.String(bigHash))
	return resources, bigHash, nil
}


func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)    
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, RENDERED_TEMPLATE_FILE)
	exists, err := fileExists(path)  
  
	if err != nil {
		return err
	}
  
	if !exists {
		resourceData.SetId("")
    return nil
	}
  
	return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)  
	resourceHashes, _, err := buildTemplateReference(providerMeta)
	
  if err != nil {
		return err
	}
  
	var templateData map[string]interface{} = map[string]interface{}{}
	var resources = map[string]interface{}{}
	templateData["Resources"] = resources
  
	for name, _ := range resourceHashes {
		properties, _, err := readResource(name, providerMeta)
  
		if err != nil {
			return err
		}
  
		data := map[string]interface{}{}
		// data["Type"] = entry.CfnType
		data["Type"] = "COMING SOON"
  
		resources[name] = data
    delete(properties, "logical_id")
  
    data["Properties"] = deSnake(properties)
	}
  
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, RENDERED_TEMPLATE_FILE)
	hash, err := writeFile(path, templateData)
  
	if err != nil {
		return err
	}
  
	resourceData.SetId(hash)  
	return nil
}


func TemplateUpdate(resourceData *schema.ResourceData, meta interface{}) error {

  // should never happen since customized diff
  panic("not updatable")
  
  return nil
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
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	_, bigHash, err := buildTemplateReference(providerMeta)
  
  if err != nil {
    return err
  }
  
  resourceDiff.SetNew("hash", bigHash)
	return nil
}
