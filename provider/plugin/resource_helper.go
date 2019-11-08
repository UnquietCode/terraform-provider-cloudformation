package plugin

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Thing struct {
  resources   map[string]string
}

// func getTemplate(provider *Provider) {
func getTemplate() Thing {
  // return provider.resources["template"]
  
  container := Thing{
    resources: make(map[string]string),
  }
  
  return container
}


func ResourceCreate(resourceName string, resourceData *schema.ResourceData, meta interface{}) error {
	// template := getTemplate()

  
  // if _, ok := template.resources[resourceName]; !ok {
  //   return nil
  // }
  // template.resources[resourceName] = resourceData

  return nil
}

func ResourceRead(resourceName string, resourceData *schema.ResourceData, meta interface{}) error {
	// template := getTemplate()
  
  // if !template.resources[resourceName] {
  //   return "Err"
  // }
  // 
  // return template[resourceName]

  return nil
}

func ResourceUpdate(resourceName string, resourceData *schema.ResourceData, meta interface{}) error {
	// template := getTemplate()
  
  // if !template.resources[resourceName] {
  //   return "Err"
  // }
  return nil
}

func ResourceDelete(resourceName string, resourceData *schema.ResourceData, meta interface{}) error {
	// template := getTemplate()

  // if !template.resources[resourceName] {
  //   return "Err"
  // }
  // delete(template.resources, resourceName)

  return nil
}