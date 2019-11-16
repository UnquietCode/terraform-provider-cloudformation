package plugin


import (
  "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"encoding/json"
)

func mapToJson(data interface{}, indent bool) ([]byte, error) {
  if indent {
	  return json.MarshalIndent(data, "", "  ") 
  } else {
    return json.Marshal(data)
  }
}

func jsonToMap(data []byte) (map[string]interface{}, error) {
  var result map[string]interface{} = map[string]interface{}{}
  err := json.Unmarshal(data, &result)
  
  if err != nil {
    return nil, err
  }
  
  return result, nil
}


type ResourceGetter func(key string, data *schema.ResourceData) (interface{}, bool)

func normalGetter(key string, data *schema.ResourceData) (interface{}, bool) {
  if value, ok := data.GetOk(key); ok {
    return value, true
  }
  return nil, false
}

func changedOnlyGetter(key string, data *schema.ResourceData) (interface{}, bool) {
  if ok := data.HasChange(key); ok {
    return data.Get(key), true
  }
  return nil, false
}

func convertValue(prefix string, resourceSchema *schema.Schema, value interface{}, resourceData *schema.ResourceData, propertyNames map[string]string, getter ResourceGetter) interface{} {
  switch resourceSchema.Type {
    case schema.TypeBool:
        fallthrough 
    case schema.TypeInt:
        fallthrough 
    case schema.TypeFloat:
        fallthrough 
    case schema.TypeString:
      return value
    
    case schema.TypeSet:
      fallthrough
    
    case schema.TypeList:
      var listIn []interface{} = value.([]interface{})
      var listOut []interface{} = make([]interface{}, 0, len(listIn))

      for idx, _ := range listIn {
        realKey := fmt.Sprintf("%s%d", prefix, idx)
        
        if listValue, ok := getter(realKey, resourceData); ok {  
          nestedPrefix := realKey + "."
        
          switch elementType := resourceSchema.Elem.(type) { 
            case *schema.Schema:
              listOut = append(listOut, convertValue(nestedPrefix, elementType, listValue, resourceData, propertyNames, getter))
            case *schema.Resource:
              listOut = append(listOut, convertResourceToMap(nestedPrefix, elementType, resourceData, propertyNames, getter))
            default:
              panic("invalid element type")
          }
        }
      }
      
      return listOut
    
    case schema.TypeMap:
      var mapIn map[string]interface{} = value.(map[string]interface{})
      var mapOut map[string]interface{} = make(map[string]interface{})
      
      for name, _ := range mapIn {
        realKey := prefix + name
        
        if mapValue, ok := getter(realKey, resourceData); ok {
          nestedPrefix := realKey + "."
          mapOut[name] = convertValue(nestedPrefix, resourceSchema.Elem.(*schema.Schema), mapValue, resourceData, propertyNames, getter)
        }
      }
      
      return mapOut
    
    default:
      panic("invalid type")
  }
}


func convertAndMergeResourceToMap(prefix string, resource *schema.Resource, resourceData *schema.ResourceData, data map[string]interface{}, propertyNames map[string]string, getter ResourceGetter) {
  for name, attribute := range resource.Schema {
		var realKey string = prefix + name
    
		if value, ok := getter(realKey, resourceData); ok {
      nestedPrefix := realKey + "."
      var pName string
      
      if _, ok := propertyNames[name]; ok {
        pName = propertyNames[name]
      } else {
        pName = name
	  	}
      
      data[pName] = convertValue(nestedPrefix, attribute, value, resourceData, propertyNames, getter)
		}
  }
}

func convertResourceToMap(prefix string, resource *schema.Resource, resourceData *schema.ResourceData, propertyNames map[string]string, getter ResourceGetter) map[string]interface{} {  
  data := map[string]interface{}{}
  convertAndMergeResourceToMap(prefix, resource, resourceData, data, propertyNames, getter)
  return data
}