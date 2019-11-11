package plugin

import (
	"errors"
	"log"
	"os"
	"fmt"
	"sort"
	"time"
	"bufio"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const RENDERD_TEMPLATE_NAME string = "template.rendered"


func readResources(meta ProviderMetadata) (map[string]TemplateEntry, error) {
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
	
	var mapD map[string]TemplateEntry = map[string]TemplateEntry{}
	
	// TODO sorting?
	// lines = sort.Strings(lines)
	
	for _, line := range lines {
		entryData, err := jsonToMap([]byte(line))
		
		if err != nil {
			return nil, err
		}
		
		var entry TemplateEntry = TemplateEntry{
			logicalId: entryData["id"].(string),
			cfnType: entryData["type"].(string),
			hash: entryData["hash"].(string),
		}
		
		mapD[entry.logicalId] = entry
	}
	
  return mapD, scanner.Err()
}


func buildTemplateReference(resourceData *schema.ResourceData, meta interface{}) ([]TemplateEntry, string, error) {
	var resourceHashes map[string]TemplateEntry
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
	
	var entries []TemplateEntry
	var bigHash string = ""
	
	for _, key := range keys {
		entry := resourceHashes[key]
		bigHash += entry.hash
		entries = append(entries, entry)
	}
	
	bigHash = strconv.Itoa(hashcode.String(bigHash))
	return entries, bigHash, nil
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
		// Delay: 1 * time.Second,
  }
	
	_, err = createStateConf.WaitForState()
	
	if err != nil {
		return err
	}
	
	var templateData map[string]interface{} = map[string]interface{}{}
	var resources = map[string]interface{}{}
	templateData["Resources"] = resources
	
	for _, entry := range refs {
		properties, _, err := readFile(entry.logicalId, providerMeta)
		
		if err != nil {
			return err
		}
		
		data := map[string]interface{}{}
		data["Type"] = entry.cfnType		
		data["Properties"] = properties
		
		resources[entry.logicalId] = data
	}
	
	hashcode, err := writeFile(RENDERD_TEMPLATE_NAME, templateData, providerMeta)
	
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
	err := removeFile(RENDERD_TEMPLATE_NAME, meta.(ProviderMetadata))
	
	if err != nil {
		return err
	}
	
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
