package plugin

import (
	"fmt"
  "crypto/md5"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// --------------------------	------------------------------------------------

func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	// var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
    
  var data map[string]interface{} = convertResourceToMap("", resourceSchema, resourceData, normalGetter)
  delete(data, "logical_id")
  properties := deSnake(data)

  resourceDataContainer := map[string]interface{}{
    "Type": resourceType,
    "Properties": properties,
  }
  
  fullData := map[string]interface{}{
    "Resources": map[string]interface{}{
      logicalId: resourceDataContainer,
    },
  }
  
  fullDataBytes, err := mapToJson(fullData, true)
  if err != nil { return err }
  
  resourceData.Set("json", string(fullDataBytes))
  
  var hashcode string = fmt.Sprintf("%x", md5.Sum(fullDataBytes))
  resourceData.SetId(hashcode)

  return nil
}