package plugin


import (
  "fmt"
  "log"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"encoding/json"
)

func mapToJson(data interface{}, indent bool) ([]byte, error) {
  log.Printf("")
  
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

// func changedOnlyGetter(key string, data *schema.ResourceData) (interface{}, bool) {
//   if ok := data.HasChange(key); ok {
//     return data.Get(key), true
//   }
//   return nil, false
// }

func convertValue(prefix string, resourceSchema *schema.Schema, value interface{}, resourceData *schema.ResourceData, getter ResourceGetter) interface{} {
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
      var setIn *schema.Set = value.(*schema.Set)
      var listIn []interface{} = setIn.List()
      var listOut []interface{} = make([]interface{}, 0, len(listIn))

      for _, oValue := range listIn {
        hashId := setIn.F(oValue)
        realKey := fmt.Sprintf("%s%d", prefix, hashId)
        
        if listValue, ok := getter(realKey, resourceData); ok {  
          nestedPrefix := realKey + "."
          
          switch elementType := resourceSchema.Elem.(type) { 
            case *schema.Schema:
              listOut = append(listOut, convertValue(nestedPrefix, elementType, listValue, resourceData, getter))
            case *schema.Resource:
              listOut = append(listOut, convertResourceToMap(nestedPrefix, elementType, resourceData, getter))
            default:
              panic("invalid element type")
          }
        }
      }
      
      if resourceSchema.MaxItems == 1 {
        return listOut[0]
      } else {
        return listOut
      }
    
    case schema.TypeList:
      var listIn []interface{} = value.([]interface{})
      var listOut []interface{} = make([]interface{}, 0, len(listIn))

      for idx, _ := range listIn {
        realKey := fmt.Sprintf("%s%d", prefix, idx)
        
        if listValue, ok := getter(realKey, resourceData); ok {  
          nestedPrefix := realKey + "."
        
          switch elementType := resourceSchema.Elem.(type) { 
            case *schema.Schema:
              listOut = append(listOut, convertValue(nestedPrefix, elementType, listValue, resourceData, getter))
            case *schema.Resource:
              listOut = append(listOut, convertResourceToMap(nestedPrefix, elementType, resourceData, getter))
            default:
              panic("invalid element type")
          }
        }
      }
      
      return listOut
    
    case schema.TypeMap:
      var mapIn map[string]interface{} = value.(map[string]interface{})
      var mapOut map[string]interface{} = map[string]interface{}{}
      
      for name, _ := range mapIn {
        realKey := prefix + name
        
        if mapValue, ok := getter(realKey, resourceData); ok {
          nestedPrefix := realKey + "."
          mapOut[name] = convertValue(nestedPrefix, resourceSchema.Elem.(*schema.Schema), mapValue, resourceData, getter)
        }
      }
      
      return mapOut
    
    default:
      panic("invalid type")
  }
}

func mergeJsonMaps(mapA map[string]interface{}, mapB map[string]interface{}) {  
  for nameB, valueB := range mapB {

    // if key is present in A
    if valueA, ok := mapA[nameB]; ok {
      
      // if it is a map, merge
      if vMap, ok := valueB.(map[string]interface{}); ok {
        mergeJsonMaps(valueA.(map[string]interface{}), vMap)
        continue
      }
    
    // else, clobber
    } else {
      mapA[nameB] = valueB
    }
  }
}


func convertAndMergeResourceToMap(prefix string, resource *schema.Resource, resourceData *schema.ResourceData, data map[string]interface{}, getter ResourceGetter) {
  newData := map[string]interface{}{}
  
  for name, attribute := range resource.Schema {
    var realKey string = prefix + name
    
    if value, ok := getter(realKey, resourceData); ok {
      nestedPrefix := realKey + "."
      newData[name] = convertValue(nestedPrefix, attribute, value, resourceData, getter)
    }
  }

  mergeJsonMaps(data, newData)
}


func convertResourceToMap(prefix string, resource *schema.Resource, resourceData *schema.ResourceData, getter ResourceGetter) map[string]interface{} {  
  data := map[string]interface{}{}
  convertAndMergeResourceToMap(prefix, resource, resourceData, data, getter)
  return data
}