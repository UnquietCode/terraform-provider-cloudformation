package plugin

// TODO can we do hashcode of a file without reading it in?
// TODO resort imports

import (
	"errors"
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"sort"
	"bufio"
	"time"
	"encoding/json"
	"strings"
	"strconv"
	"crypto/md5"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	mutex "github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

type ProviderMetadata struct {
   workdir string
	 // existingResources *[]string
	 resourceCounter *int
	 mutex *mutex.MutexKV
	 
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


func readFile(id string, meta ProviderMetadata) (map[string]interface{}, string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	exists, err := fileExists(path);
	
	if err != nil {
		return nil, EMPTY, err
	}
	
	if !exists {
		return nil, EMPTY, nil
	}
	
	// try to read the file
	rawData, hashcode, err := readAndHashFile(path)
	
	if err != nil {
		return nil, EMPTY, err
	}
	
	// turn the data into schema data
	var data map[string]interface{}

  if err := json.Unmarshal(rawData, &data); err != nil {
    return nil, EMPTY, err
  }
	
	return data, hashcode, nil
}


func readResources(meta ProviderMetadata) (map[string]string, error) {
	var path string = fmt.Sprintf("%s/template.data.json", meta.workdir)
  file, err := os.Open(path)
  
	if err != nil {
      return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  
	for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
	
	err = scanner.Err()
	
	if err != nil {
		return nil, err
	}
	
	var mapD map[string]string = map[string]string{}
	// lines = sort.Strings(lines)
	
	for _, line := range lines {
		parts := strings.Split(line, " ")
		mapD[parts[0]] = parts[1]
	}
	
  return mapD, scanner.Err()
}


func removeFile(id string, meta ProviderMetadata) error {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	os.Remove(path)
	
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


func handleExistingResouce(meta ProviderMetadata, id string, hash string) error {
	meta.mutex.Lock(LOCK_TEMPLATE_REFERENCE)
	defer meta.mutex.Unlock(LOCK_TEMPLATE_REFERENCE)
	
	var file *os.File = nil
	var path string = fmt.Sprintf("%s/template.data.json", meta.workdir)
	
	// if this is the first run, create the file (or overwrite it)
	// if len(*meta.existingResources) == 0 {
	if (*meta.newIndex == true) {
		f, err := os.Create(path)
		
		if err != nil {
			return err
		}
		
		file = f
		*meta.newIndex = false
	}
	
	// open the file if it wasn't created
	if file == nil {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		
		if err != nil {
			return err
		}
		
		file = f
	}
	
	// append to the file
	file.WriteString(id+" "+hash+"\n")
	file.Close()
	
	// append to list
	// *meta.existingResources = append(*meta.existingResources, id)	
	return nil
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


// --------------------------	------------------------------------------------

func ResourceExists(resourceData *schema.ResourceData, meta interface{}) (bool, error) {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, logicalId)

	log.Printf("-------------- exists %s", logicalId)

	exists, err := fileExists(path);	

	if err != nil {
		return false, err
	}
	
	(*providerMeta.exists)[logicalId] = exists


	// if exists {
		// rawData, hashcode, err := readAndHashFile(path)
		// handleExistingResouce(providerMeta, logicalId, hash)
	// }

  return exists, nil
}


func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	
	data, hash, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	// doesn't exist
	if data == nil && hash == EMPTY {
		// handleExistingResouce(providerMeta, logicalId, "")
		resourceData.SetId("")
		return nil
	}
	
	for name, value := range data {
		resourceData.Set(name, value)
	}
	
	resourceData.SetId(hash)
	// handleExistingResouce(providerMeta, logicalId, hash)
	incrementResourceCounter(providerMeta)
	
  return nil
}


func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var data map[string]interface{} = make(map[string]interface{})

	for name := range resourceSchema.Schema {
		if value, ok := resourceData.GetOk(name); ok {
	  	data[name] = value
		}
	}
	
	var logicalId string = resourceData.Get("logical_id").(string)
	hashcode, err := writeFile(logicalId, data, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId(hashcode)
	incrementResourceCounter(providerMeta)
	
  return nil
}


func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	data, _, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	for name := range resourceSchema.Schema {
		if resourceData.HasChange(name) {
			_, value := resourceData.GetChange(name)
			data[name] = value
		}
	}
	
	hashcode, err := writeFile(logicalId, data, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId(hashcode)
	incrementResourceCounter(providerMeta)

  return nil
}


func ResourceDelete(resourceType string, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	err := removeFile(logicalId, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId("")
	incrementResourceCounter(providerMeta)
	
  return nil
}


func ResourceCustomizeDiff(resourceDiff *schema.ResourceDiff, meta interface{}) error {
	// log.Printf("%s", resourceDiff)
	// log.Printf("%s", meta.(ProviderMetadata).existingResources)
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	
	logicalId := resourceDiff.Get("logical_id").(string)
	
	if _, ok := (*providerMeta.exists)[logicalId]; ok {
    return errors.New("exists, duplicate logical id "+logicalId)
	}

	if _, ok := (*providerMeta.diffed)[logicalId]; ok {
    return errors.New("diffed, duplicate logical id "+logicalId)
	}
	
	(*providerMeta.diffed)[logicalId] = true
	
	log.Printf("customizeResourceDiff %s +++++++++++++++++++++++++++++++++++++++++++++++++", logicalId)

	handleExistingResouce(providerMeta, logicalId, "---")
	
	// if id, ok := resourceDiff.GetOk("id"); ok {
	// 	handleExistingResouce(meta.(ProviderMetadata), logical_id.(string), id.(string))
	// } else {
	// 	handleExistingResouce(meta.(ProviderMetadata), logical_id.(string), "")
	// 
	// }
	
	return nil
}

// --------------------------------------------------------------------------

func buildTemplateReference(resourceData *schema.ResourceData, meta interface{}) ([]string, string, error) {
	var resourceHashes map[string]string
	var err error
	resourceHashes, err = readResources(meta.(ProviderMetadata))
	
	if err != nil {
		return nil, EMPTY, err
	}
	
	var keys []string
	
	for key, _ := range resourceHashes {
    keys = append(keys, key)
	}
	
	sort.Strings(keys)
	
	var bigHash string = ""
	
	for _, key := range keys {
		hash := resourceHashes[key]
		bigHash += hash
	}
	
	bigHash = strconv.Itoa(hashcode.String(bigHash))
	return keys, bigHash, nil
}


func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
	var path string = fmt.Sprintf("%s/template.rendered.json", meta.(ProviderMetadata).workdir)	
	exists, err := fileExists(path);
	
	if err != nil {
		return err
	}
	
	if !exists {
		resourceData.SetId("")
	} else {
		_, hashcode, err := readAndHashFile(path)
	
		if err != nil {
			return err
		}
		
		resourceData.SetId(hashcode)
	}
	
	return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	refs, _, err := buildTemplateReference(resourceData, meta)
	var refCount int = len(refs)
	
	if err != nil {
		return err
	}
	
	createStateConf := &resource.StateChangeConf{
    Pending: []string{
      STATE_WAIT,
    },
    Target: []string{
      STATE_DONE,
    },
    Refresh: func() (interface{}, string, error) {
			var counter int = readCounter(providerMeta)
			log.Printf("waiting for template refs to converge %d %d", refCount, counter)
			
      if counter == refCount {
				return STATE_DONE, STATE_DONE, nil
			}
			
      if counter >= refCount {
				return nil, EMPTY, errors.New("too many refs")
			}

			return nil, STATE_WAIT, nil
    },
		Timeout: 10 * time.Minute,
		Delay: 1 * time.Second,
  }
	
	_, err = createStateConf.WaitForState()
	
	if err != nil {
		return err
	}
	
	var templateData map[string]interface{} = map[string]interface{}{}
	templateData["Resources"] = map[string]interface{}{}
	
	for _, ref := range refs {
		data, _, err := readFile(ref, providerMeta)
		
		if err != nil {
			return err
		}
		
		templateData["Resources"].(map[string]interface{})[ref] = data
	}
	
	hashcode, err := writeFile("template.rendered", templateData, providerMeta)
	
	if err != nil {
		return err
	}
	
	resourceData.SetId(hashcode)
	return nil
}


func TemplateUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	// refs, hash, err := buildTemplateReference(resourceData, meta)
	
	// resourceData.SetId(bigHash)
  return nil
}


func TemplateDelete(resourceData *schema.ResourceData, meta interface{}) error {
	resourceData.SetId("")

  return nil
}

func TemplateCustomizeDiff(resourceDiff *schema.ResourceDiff, meta interface{}) error {
	log.Printf("customizeDiff +++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Printf("%s", resourceDiff)
	// log.Printf("%s", meta.(ProviderMetadata).resourceHashes)
	resourceDiff.SetNewComputed("id")
	return nil
}

// ---------------------------------------------------------------------------

func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	var counter int = 0
	var newIndex bool = true
	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		resourceCounter: &counter,
		// existingResources: &[]string{},
		mutex: mutex.NewMutexKV(),
		
		newIndex: &newIndex,
		exists: &map[string]bool{},
		diffed: &map[string]bool{},
  }
	
	return meta, nil
}