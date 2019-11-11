package plugin

import (
	"errors"
	"strings"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ResourceMeat struct {
  Type string
	Properties map[string]interface{}
}

type TemplateData struct {
  resources map[string]ResourceMeat
}

type ProviderMetadata struct {
   template TemplateData
	 counters map[string]int
}


func newTemplate() TemplateData {
  // return provider.resources["template"]
  
  container := TemplateData{
    resources: make(map[string]ResourceMeat),
  }
  
  return container
}

func getTemplate(meta interface{}) TemplateData {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	return providerMeta.template
}

func getResource(id string, cfnType string, template TemplateData) ResourceMeat {
	if val, ok := template.resources[id]; ok {
  	return val
	} else {
		resource := ResourceMeat{
			Type: cfnType,
			Properties: make(map[string]interface{}),
		}
		template.resources[id] = resource
		return resource
	}
}

func convertFromTerraform(data interface{}) interface{} {
	return nil
}

func convertToTerraform(data interface{}) interface{} {
	return nil
}

func getId(resourceType string, resourceData *schema.ResourceData, meta ProviderMetadata) string {
	if value, ok := resourceData.GetOk("logical_id"); ok {
  	return value.(string)
	
	// make one up
	} else {
	 	var parts = strings.Split(resourceType, "::")
		var counterName = parts[1] + parts[2]
		
		if _, ok := meta.counters[counterName]; !ok {
			meta.counters[counterName] = 1
		}
		
		var counter int = meta.counters[counterName]
		meta.counters[counterName] = counter + 1
		
		return counterName + strconv.Itoa(counter)
	}
}


func ResourceCreate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = getId(resourceType, resourceData, meta.(ProviderMetadata))
	var template TemplateData = getTemplate(meta)
  
  if _, ok := template.resources[logicalId]; ok {
    return errors.New("already exists?")
  } else {
		var resource = getResource(logicalId, resourceType, template)
				
		resourceData.Partial(true)
		
		for name := range resourceSchema.Schema {
			if value, ok := resourceData.GetOk(name); ok {
		  	resource.Properties[name] = convertFromTerraform(value)
				resourceData.SetPartial(name)
			}
		}
		
		resourceData.SetId(logicalId)
		resourceData.Partial(false)
	}

  return nil
}

func ResourceRead(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = getId(resourceType, resourceData, meta.(ProviderMetadata))
	var template TemplateData = getTemplate(meta)
	var resource = getResource(logicalId, resourceType, template)
	
	for name := range resource.Properties {
  	resourceData.Set(name, convertToTerraform(resource.Properties[name]))
	}
	
  return nil
}

func ResourceUpdate(resourceType string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = getId(resourceType, resourceData, meta.(ProviderMetadata))
	var template TemplateData = getTemplate(meta)
	var resource = getResource(logicalId, resourceType, template)
	
	resourceData.Partial(true)
	
	for name := range resourceSchema.Schema {
		if resourceData.HasChange(name) {
			_, value := resourceData.GetChange(name)
			resource.Properties[name] = convertFromTerraform(value)
		}
	}

	resourceData.Partial(false)
	
  return nil
}

func ResourceDelete(resourceType string, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = getId(resourceType, resourceData, meta.(ProviderMetadata))
	var template TemplateData = getTemplate(meta)
	
	if _, ok := template.resources[logicalId]; ok {
  	delete(template.resources, logicalId)
	} else {
		return errors.New("resource does not exist?")
	}

  return nil
}


func ProviderConfigure(resourceData *schema.ResourceData) (interface{}, error) {
	
	// check for stack in CF
		// if it exists, get the template data
		// if it doesn't exist, start a new template
	var template TemplateData = newTemplate()
	
	meta := ProviderMetadata{
    template: template,
		counters: make(map[string]int),
  }
	
	return meta, nil
}