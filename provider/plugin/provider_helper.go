package plugin


import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"crypto/md5"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	mutex "github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

type ProviderMetadata struct {
   workdir string
	 mutex *mutex.MutexKV
	 resourceCounter *int
	 
	 newIndex *bool
	 exists *map[string]bool
	 diffed *map[string]bool
}

const EMPTY string = "\xff"
const LOCK_TEMPLATE_REFERENCE string = "TEMPLATE_REFERENCE"
const LOCK_RESOURCE_COUNTER string = "RESOURCE_COUNTER"
const STATE_WAIT string = "WAIT"
const STATE_DONE string = "DONE"


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


func writeFile(name string, data map[string]interface{}, meta ProviderMetadata) (string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, name)

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


func incrementResourceCounter(meta ProviderMetadata) {
	meta.mutex.Lock(LOCK_RESOURCE_COUNTER)
	*meta.resourceCounter += 1
	meta.mutex.Unlock(LOCK_RESOURCE_COUNTER)
}


func readCounter(meta ProviderMetadata) int {
	meta.mutex.Lock(LOCK_RESOURCE_COUNTER)
	defer meta.mutex.Unlock(LOCK_RESOURCE_COUNTER)
	return *meta.resourceCounter
}


func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	var counter int = 0
	var newIndex bool = true
	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		mutex: mutex.NewMutexKV(),
		resourceCounter: &counter,
		
		newIndex: &newIndex,
		exists: &map[string]bool{},
		diffed: &map[string]bool{},
  }
	
	return meta, nil
}