package plugin


import (
  "strings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	mutex "github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

type ProviderMetadata struct {
   workdir string
	 mutex *mutex.MutexKV
	 newIndex *map[string]bool
	 exists *map[string]bool
	 diffed *map[string]bool
}

const EMPTY string = "\xff"
const LOCK_RESOURCE_READ_WRITE string = "RESOURCE_READ_WRITE"
const LOCK_RESOURCE_READ_COUNT string = "RESOURCE_READ_COUNT"

const TEMPLATE_INDEX_FILE string = "template.data"
const TEMPLATE_COUNTER_FILE string = "template.counts"
const RENDERED_TEMPLATE_FILE string = "template.rendered"

const STATE_WAIT string = "WAIT"
const STATE_DONE string = "DONE"


type TemplateEntry struct {
  CfnType string `json:"type"`
  Hash string `json:"hash"`
}


func allWithout(item string, items...string) []string {
  newArray := []string{}
  
  for _, x := range items {
    if x != item {
      newArray = append(newArray, x)
    }
  }
  
  return newArray
}


func ensureArrays(mapData map[string]interface{}, keys...string) {
  for _, key := range keys {
    if _, ok := mapData[key]; !ok {
      mapData[key] = []interface{}{}
    }
  }
}

func addString(item string, addArray []interface{}) []interface{} {
  var alreadyExists bool = false
  
  for _, v := range addArray {
    if v.(string) == item {
      alreadyExists = true
      break
    }
  }
  
  if !alreadyExists {
    addArray = append(addArray, item)
  }
  
  return addArray
}

func removeString(item string, removeArray []interface{}) []interface{} {
  var newRemoved []interface{} = []interface{}{}
  
  for _, v := range removeArray {
    if v.(string) != item {
      newRemoved = append(newRemoved, v.(string))
    }
  }
  
  return newRemoved
}

func arrayContainsString(array []string, item string) bool {
  for _, v := range array {
    if v == item {
      return true
    }
  }
  return false
}


func arrayContainsItem(array interface{}, item interface{}) bool {
  for _, v := range array.([]interface{}) {
    if v == item {
      return true
    }
  }
  return false
}


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
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		mutex: mutex.NewMutexKV(),
		newIndex: &map[string]bool{
      TEMPLATE_COUNTER_FILE: true,
      TEMPLATE_INDEX_FILE: true,
    },
		exists: &map[string]bool{},
		diffed: &map[string]bool{},
  }
	
	return meta, nil
}