package plugin

import (
	"fmt"
  "errors"
  "crypto/md5"  
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
  // var providerMeta ProviderMetadata = meta.(ProviderMetadata)
  
  // merge all JSON together
  var fullData map[string]interface{} = map[string]interface{}{
    "Description": resourceData.Get("description").(string),
    "Resources": map[string]interface{}{},
  }
  allResources := fullData["Resources"].(map[string]interface{})
  
  for _, item := range resourceData.Get("resources").([]interface{}) {
    data, err := jsonToMap([]byte(item.(string)))
    if err != nil { return err }
    
    for n,v := range data["Resources"].(map[string]interface{}) {
      if _, ok := allResources[n]; ok {
        return errors.New("duplicate logical ID '"+n+"'")
      }
      allResources[n] = v
    }
  }
  
  fullDataBytes, err := mapToJson(fullData, true)
  if err != nil { return err }
  
  resourceData.Set("json", string(fullDataBytes))

  var hashcode string = fmt.Sprintf("%x", md5.Sum(fullDataBytes))
  resourceData.SetId(hashcode)  
  
  return nil
}
