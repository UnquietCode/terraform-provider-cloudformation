package plugin

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type TemplateData struct {
  resources map[string]map[string]interface{}
}

type ProviderMetadata struct {
   template TemplateData
}


func newTemplate() TemplateData {
  // return provider.resources["template"]
  
  container := TemplateData{
    resources: make(map[string]map[string]interface{}),
  }
  
  return container
}

func getTemplate(meta interface{}) TemplateData {
	var providerMeta ProviderMetadata = meta.(ProviderMetadata)
	return providerMeta.template
}

func getResource(id string, template TemplateData) map[string]interface{} {
	if val, ok := template.resources[id]; ok {
  	return val
	} else {
		var resource = make(map[string]interface{})
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

func deSnake(name string) string {
		return name
		
    // dname = ""
    // dname += name[0].upper()
    // dname += name[1:]
		// 
    // # after an underscore
    // repl_name = re.sub(
    //     string=dname,
    //     pattern=r'_([^_])([^_]*)',
    //     repl=r'$\2',
    // )
		// 
    // dname = repl_name + ""
		// 
    // for c, idx in enumerate(repl_name):
    //     if c == '$':
    //         dname[idx] = repl_name[idx].upper()
		// 
    // return dname
}


func ResourceCreate(resourceName string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = deSnake(resourceName)
	var template TemplateData = getTemplate(meta)
  
  if _, ok := template.resources[logicalId]; ok {
    return errors.New("already exists?")
  } else {
		var resource = make(map[string]interface{})
  	template.resources[resourceName] = resource
		
		for name := range resourceSchema.Schema {
	  	resource[name] = convertFromTerraform(resourceData.Get(name))
		}
	}

  return nil
}

func ResourceRead(resourceName string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = deSnake(resourceName)
	var template TemplateData = getTemplate(meta)
	var resource = getResource(logicalId, template)
	
	for name := range resourceSchema.Schema {
  	resourceData.Set(name, convertToTerraform(resource[name]))
	}
	
  return nil
}

func ResourceUpdate(resourceName string, resourceSchema *schema.Resource, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = deSnake(resourceName)
	var template TemplateData = getTemplate(meta)
	var resource = getResource(logicalId, template)
	
	resourceData.Partial(true)
	
	for name := range resourceSchema.Schema {
		if resourceData.HasChange(name) {
			_, value := resourceData.GetChange(name)
			resource[name] = convertFromTerraform(value)
		}
	}

	resourceData.Partial(false)
	
  return nil
}

func ResourceDelete(resourceName string, resourceData *schema.ResourceData, meta interface{}) error {
	var logicalId string = deSnake(resourceName)
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
  }
	
	return meta, nil
}