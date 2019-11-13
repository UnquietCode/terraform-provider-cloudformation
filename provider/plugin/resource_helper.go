package plugin

// TODO can we do hashcode of a file without reading it in?

import (
	"errors"
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"sort"
	"bufio"
	"encoding/json"
	"strings"
	"sync/atomic"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type ResourceMeat struct {
  Type string
	Properties map[string]interface{}
}

type ProviderMetadata struct {
   workdir string
	 existingResources []string
	 templateTargets []string
	 resourceHashes map[string]interface{}
	 resourceCounter int32
}

const EMPTY string = "\xff"


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
	rawData, err := ioutil.ReadFile(path)
	var hashcode string = fmt.Sprintf("%s", hashcode.String(string(rawData)))
	
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
	
	for line := range lines {
		parts := strings.Split(lines[line], " ")
		mapD[parts[0]] = parts[1]
	}
	
  return mapD, scanner.Err()
}


func removeFile(id string, meta ProviderMetadata) error {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	os.Remove(path)
	
	return nil
}


func writeFile(id string, data map[string]interface{}, meta ProviderMetadata) (string, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)

	file, _ := os.Create(path)
	defer file.Close()
	
	rawData, err := json.Marshal(data)
	
	if err != nil {
		return "", err
	}
	
 	file.Write(rawData)
	
	var hashcode string = fmt.Sprintf("%s", hashcode.String(string(rawData)))
	return hashcode, nil
}


func readAllFiles(path string) string {
  files, _ := ioutil.ReadDir(path)
	var data string = ""
	
	for _, f := range files {
		log.Printf(f.Name())
    buf, _ := ioutil.ReadFile(path+"/"+f.Name())
		data += string(buf)
  }
	
	return data
}

func readAll(resourceData *schema.ResourceData, meta interface{}) error {
	
	log.Printf("here again +++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Printf("hey %s", template.resources)
	
	// resourceData.SetId("")

  return nil
}

func handleExistingResouce(meta ProviderMetadata, id string, hash string) error {
	// TODO mutex
	var file *os.File = nil
	var path string = fmt.Sprintf("%s/template.data.json", meta.workdir)
	
	// if this is the first run, create the file (or overwrite it)
	if len(meta.existingResources) == 0 {
		f, err := os.Create(path)
		
		if err != nil {
			return err
		}
		
		file = f
	}
	
	// open the file if it wasn't
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
	
	// append to the list
	meta.existingResources = append(meta.existingResources, id)	
	return nil
}

// --------------------------	------------------------------------------------

func ResourceExists(resourceData *schema.ResourceData, meta interface{}) (bool, error) {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	var path string = fmt.Sprintf("%s/%s.json", providerMeta.workdir, logicalId)
	
	log.Printf("exists +++++++++++++++++++++++++++++++++++++++++++++++++")
	
	exists, err := fileExists(path);	
	
	if err != nil {
		return false, err
	}
	
  return exists, nil
}

func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	data, hash, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	for name := range data {
		resourceData.Set(name, data[name])
	}
	
	resourceData.SetId(hash)
	handleExistingResouce(providerMeta, logicalId, hash)
	atomic.AddInt32(&providerMeta.resourceCounter, 1)
	
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
	atomic.AddInt32(&providerMeta.resourceCounter, 1)
  return nil
}


func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
  // var data map[string]interface{} = map[string]interface{}{}
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
	atomic.AddInt32(&providerMeta.resourceCounter, 1)
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
	atomic.AddInt32(&providerMeta.resourceCounter, 1)
  return nil
}

func ResourceCustomizeDiff(resourceDiff *schema.ResourceDiff, meta interface{}) error {
	// resourceDiff.SetNewComputed("last_
	return nil
}

// --------------------------------------------------------------------------

func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
	// hash, _, err := buildTemplateReference(resourceData, meta)
	// 
	// if err != nil {
	// 	return err
	// }
	// 
	// resourceData.SetId(hash)
	return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	
	refs, hash, err := buildTemplateReference(resourceData, meta)
	
	if err != nil {
		return err
	}
	
	createStateConf := &resource.StateChangeConf{
      Pending: []string{
        "WAIT",
      },
      Target: []string{
        "DONE",
      },
      Refresh: func() (interface{}, string, error) {
				var counter int32 = atomic.LoadInt32(&providerMeta.resourceCounter)
				var refs32 int32 = int32(refs)
				
        if counter == refs32 {
					return nil, "DONE", nil
				}
				
        if counter >= refs32 {
					return nil, EMPTY, errors.New("too many refs")
				}

				return nil, "WAIT", nil
      },
  }
	createStateConf.WaitForState()
	
	resourceData.SetId(hash)
	return nil
}


func buildTemplateReference(resourceData *schema.ResourceData, meta interface{}) (int, string, error) {
	var resourceHashes map[string]string
	var err error
	resourceHashes, err = readResources(meta.(ProviderMetadata))
	
	if err != nil {
		return -1, EMPTY, err
	}
	
	var keys []string
	
	for k := range resourceHashes {
    keys = append(keys, k)
	}
	
	sort.Strings(keys)
	
	var bigHash string = ""
	
	for s := range keys {
		var ss string = keys[s]
		log.Printf("line %s +++++++++++++++++++++++++++++++++++++++++++++++++", ss)
		hash := resourceHashes[ss]
		bigHash += hash
	}
	
	bigHash = fmt.Sprintf("%s", hashcode.String(bigHash))
	return len(keys), bigHash, nil
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
	
	log.Printf("here +++++++++++++++++++++++++++++++++++++++++++++++++")
	// log.Printf("%s", resourceDiff)
	// log.Printf("%s", meta.(ProviderMetadata).resourceHashes)
	resourceDiff.SetNewComputed("id")
	return nil
}

// ---------------------------------------------------------------------------

func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		existingResources: []string{},
		resourceHashes: map[string]interface{}{},
		resourceCounter: 0,
  }
	
	return meta, nil
}