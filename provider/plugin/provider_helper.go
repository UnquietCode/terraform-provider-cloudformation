package plugin


import (
  "strings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ProviderMetadata struct {}

const EMPTY string = "\xff"


func replaceFunctionCalls(input interface{}) interface{} {
  if value, ok := input.(string); ok {
    if strings.HasPrefix(value, "!GetAtt") {
      property := strings.TrimSpace(value[7:])
      parts := strings.Split(property, ".")
      
      return map[string]interface{}{
        "Fn::GetAtt": []string{parts[0], parts[1]},
      }
    }
    
    if strings.HasPrefix(value, "!Ref") {
      id := strings.TrimSpace(value[4:])
      
      return map[string]interface{}{
        "Ref": id,
      }
    }
  }
  return input
}


func deSnake(input map[string]interface{}) map[string]interface{} {
  var output map[string]interface{} = map[string]interface{}{}
  
  for k,v := range(input) {
    var newString string = ""
    
    for _, part := range strings.Split(k, "_") {
      newString += strings.ToUpper(part[0:1])
      newString += part[1:]
    }
    
    if vMap, ok := v.(map[string]interface{}); ok {
      output[newString] = deSnake(vMap)
    } else if vList, ok := v.([]interface{}); ok {
      var newList []interface{} = make([]interface{}, len(vList))
      
      for i, elem := range vList {
        if eMap, ok := elem.(map[string]interface{}); ok {
          newList[i] = deSnake(eMap)
        } else {
          newList[i] = replaceFunctionCalls(elem)
        }
      }
      
      output[newString] = newList
    } else {
      output[newString] = replaceFunctionCalls(v)
    }
  }
  
  return output
}



func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	meta := ProviderMetadata{}
	
	return meta, nil
}