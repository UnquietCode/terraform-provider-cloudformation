package plugin

// TODO can we do hashcode of a file without reading it in?
// TODO resort imports

import (
	"errors"
	"log"
	"fmt"
	"sort"
	"time"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

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
