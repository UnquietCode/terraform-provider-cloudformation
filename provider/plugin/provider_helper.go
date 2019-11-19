package plugin


import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
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
	 
	 // newIndex *map[string]bool
	 exists *map[string]bool
	 diffed *map[string]bool
   
   // templateReading *bool
}

const EMPTY string = "\xff"
const LOCK_RESOURCE_READ_WRITE string = "RESOURCE_READ_WRITE"

const TEMPLATE_INDEX_FILE string = "template.data.json"
const RENDERED_TEMPLATE_FILE string = "template.rendered"

// 
// const LOCK_TEMPLATE_REFERENCE string = "TEMPLATE_REFERENCE"
// const LOCK_TEMPLATE_EXISTENCE string = "TEMPLATE_EXISTENCE"
// const LOCK_RESOURCE_COUNTER string = "RESOURCE_COUNTER"
// const LOCK_RESOURCE_READ_COUNTER string = "RESOURCE_READ_COUNTER"
// 
// const STATE_WAIT string = "WAIT"
// const STATE_DONE string = "DONE"





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
	
	rawData, err := json.MarshalIndent(data, "", "  ")
	
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
	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		mutex: mutex.NewMutexKV(),
		// resourceCounter: &counter,
    // readCounter: &readCounter,
		
		// newIndex: &map[string]bool{
      // TEMPLATE_DATA_FILE: true,
      // TEMPLATE_EXISTENCE_FILE: true,
    // },
		exists: &map[string]bool{},
		diffed: &map[string]bool{},
    // 
    // templateReading: &templateReading,
  }
	
	return meta, nil
}