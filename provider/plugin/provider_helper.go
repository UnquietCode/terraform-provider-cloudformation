package plugin


import (
	"os"
	"io/ioutil"
  "log"
	"fmt"
	"crypto/md5"
  "strings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	mutex "github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

type ProviderMetadata struct {
   workdir string
	 mutex *mutex.MutexKV
	 // resourceCounter *int
   // readCounter *int
	 
	 newIndex *map[string]bool
	 exists *map[string]bool
	 diffed *map[string]bool
   
   // templateReading *bool
}

const EMPTY string = "\xff"
const LOCK_RESOURCE_READ_WRITE string = "RESOURCE_READ_WRITE"
const LOCK_RESOURCE_READ_COUNT string = "RESOURCE_READ_COUNT"

const TEMPLATE_INDEX_FILE string = "template.data"
const TEMPLATE_COUNTER_FILE string = "template.counts"
const RENDERED_TEMPLATE_FILE string = "template.rendered"

// 
// const LOCK_TEMPLATE_REFERENCE string = "TEMPLATE_REFERENCE"
// const LOCK_TEMPLATE_EXISTENCE string = "TEMPLATE_EXISTENCE"
// const LOCK_RESOURCE_COUNTER string = "RESOURCE_COUNTER"
// const LOCK_RESOURCE_READ_COUNTER string = "RESOURCE_READ_COUNTER"
// 
const STATE_WAIT string = "WAIT"
const STATE_DONE string = "DONE"

type ChangeType = string

const (
  Unchanged ChangeType = "unchanged"
  Changed ChangeType = "changed"
  Maybe ChangeType = "maybe"
  Added ChangeType = "added"
  Deleted ChangeType = "deleted"
)

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

func ensureArraysCT(mapData map[ChangeType][]string, keys...ChangeType) {
  for _, key := range keys {
    if _, ok := mapData[key]; !ok {
      mapData[key] = []string{}
    }
  }
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

func arrayContains(array []string, item string) bool {
  for _, v := range array {
    if v == item {
      return true
    }
  }
  return false
}



// func addAndRemoveString(item string, addArray *[]interface{}, removeArrays...*[]interface{}) {
// 
//   // add
//   var alreadyExists bool = false
// 
//   for _, v := range *addArray {
//     if v.(string) == item {
//       alreadyExists = true
//       break
//     }
//   }
// 
//   if !alreadyExists {
//     addArray := append(*addArray, item)
//   }
// 
//   // remove
//   for _, removeArray := range removeArrays {
//     var newRemoved []interface{} = make([]interface{}, 0, len(*removeArray))
// 
//     for _, v := range *removeArray {
//       if v.(string) != item {
//         newRemoved = append(newRemoved, v.(string))
//       }
//     }
// 
//     *removeArray = newRemoved
//   }
// }


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
          newList[i] = elem
        }
      }
      
      output[newString] = newList
    } else {
      output[newString] = v
    }
  }
  
  return output
}



func fileExists(path string) (bool, error) {  
	_, err := os.Stat(path)
	
	if err != nil {
		if os.IsNotExist(err) {
	      return false, nil
	  }
		return false, err
	}
  
	return true, nil
}

// func writeMarker(


func readAndHashFile(path string) ([]byte, string, error) {
	rawData, err := ioutil.ReadFile(path)
	
	if err != nil {
		return nil, EMPTY, err
	}

	var hashcode string = fmt.Sprintf("%x", md5.Sum(rawData))
	return rawData, hashcode, nil
}


func removeFile(name string, meta ProviderMetadata) error {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, name)
	err := os.Remove(path)
	
	if err != nil {
		return err
	}
	
	return nil
}


func writeFile(path string, data map[string]interface{}) (string, error) {  
	file, _ := os.Create(path)
	defer file.Close()
	
	rawData, err := mapToJson(data, true)
	
	if err != nil {
		return "", err
	}
	
 	file.Write(rawData)
	
	var hashcode string = fmt.Sprintf("%x", md5.Sum(rawData))
	return hashcode, nil
}


func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	// var counter int = 0
	// var readCounter int = 0
	// var newIndex bool = true
  // var templateReading bool = false
  log.Printf("-------------------------------RECONFIGURE")
	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		mutex: mutex.NewMutexKV(),
		// resourceCounter: &counter,
    // readCounter: &readCounter,
		
		newIndex: &map[string]bool{
      TEMPLATE_COUNTER_FILE: true,
      // TEMPLATE_DATA_FILE: true,
      // TEMPLATE_EXISTENCE_FILE: true,
    },
		exists: &map[string]bool{},
		diffed: &map[string]bool{},
    // 
    // templateReading: &templateReading,
  }
	
	return meta, nil
}