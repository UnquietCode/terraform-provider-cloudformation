package plugin

// TODO can we do hashcode of a file without reading it in?

import (
	//"errors"
	"log"
	"os"
	 "io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)

type ResourceMeat struct {
  Type string
	Properties map[string]interface{}
}

type ProviderMetadata struct {
   workdir string
	 existingResources []string
	 templateTargets []string
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


func readFile(id string, meta ProviderMetadata) (map[string]interface{}, error) {
	var path string = fmt.Sprintf("%s/%s.json", meta.workdir, id)
	exists, err := fileExists(path);
	
	if err != nil {
		return nil, err
	}
	
	if !exists {
		return nil, nil
	}
	
	// try to read the file
	rawData, err := ioutil.ReadFile(path)
	
	if err != nil {
		return nil, err
	}
	
	// turn the data into schema data
	var data map[string]interface{}

  if err := json.Unmarshal([]byte(rawData), &data); err != nil {
    return nil, err
  }
	
	return data, nil
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

func handleExistingResouce(meta ProviderMetadata, id string) error {
	// TODO mutex
	var file *os.File = nil
	var path string = fmt.Sprintf("%s/template.data.json", meta.workdir)
	
	// if this is the first run, create the file
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
	file.WriteString(id+"\n")
	file.Close()
	
	// append to the list
	meta.existingResources = append(meta.existingResources, id)	
	
	return nil
}

// --------------------------	------------------------------------------------

func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	var logicalId string = resourceData.Get("logical_id").(string)
	data, err := readFile(logicalId, providerMeta)
	
	if err != nil {
		return err
	}
	
	if data != nil {
		for name := range data {
			resourceData.Set(name, data[name])
		}
		
		handleExistingResouce(providerMeta, logicalId)
	}
	
  return nil
}

func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
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
  return nil
}


func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	// var logicalId string = resourceData.Get("logical_id").(string)

	// resourceData.Partial(true)
	// 
	// for name := range resourceSchema.Schema {
	// 	if resourceData.HasChange(name) {
	// 		_, value := resourceData.GetChange(name)
	// 		resource.Properties[name] = value
	// 	}
	// }
	// 
	// resourceData.Partial(false)
	
  return nil
}


func ResourceDelete(resourceType string, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = resourceData.Get("logical_id").(string)
	err := removeFile(logicalId, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
	resourceData.SetId("")
  return nil
}

// --------------------------------------------------------------------------

func TemplateRead(resourceData *schema.ResourceData, meta interface{}) error {
	// var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	// var path string = fmt.Sprintf("%s/template.data.json", providerMeta.workdir)

	
	
	
		// f, _ := os.Create("/home/ben/Documents/Code-Projects/cfnpvd/workdir/"+id+".json")
		
		// var data string = readAllFiles("/home/ben/Documents/Code-Projects/cfnpvd/workdir")
		// resourceData.Set("output", data)
		// 
		// log.Printf("hey again +++++++++++++++++++++++++++++++++++++++++++++++++")
		// log.Printf(data)
		// 
		// var id string = fmt.Sprintf("%s", hashcode.String(data))
		// resourceData.SetId(id)
	
	// TODO mutex
	// providerMeta.templateTargets = append(providerMeta.templateTargets, "")
	

  return nil
}


func TemplateCreate(resourceData *schema.ResourceData, meta interface{}) error {

	
	// return TemplateRead(resourceData, meta)

  // return nil
	resourceData.SetId("-")
	return nil
}


func TemplateUpdate(resourceData *schema.ResourceData, meta interface{}) error {

	
	// return TemplateRead(resourceData, meta)

  return nil
}


func TemplateDelete(resourceData *schema.ResourceData, meta interface{}) error {

	// resourceData.SetId("")

  return nil
}

// ---------------------------------------------------------------------------

func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {	
	meta := ProviderMetadata{
		workdir: resourceData.Get("workdir").(string),
		existingResources: []string{},
  }
	
	return meta, nil
}