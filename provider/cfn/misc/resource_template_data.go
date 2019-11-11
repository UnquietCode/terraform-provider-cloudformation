package misc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTemplateData() *schema.Resource {
	return &schema.Resource{
		Read: 	resourceTemplateRead,
		Create: resourceTemplateCreate,
		Update: resourceTemplateUpdate,
		Delete: resourceTemplateDelete,		
		CustomizeDiff: resourceTemplateCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"output": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateRead(data, meta)
}

func resourceTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateCreate(data, meta)
}

func resourceTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateUpdate(data, meta)
}

func resourceTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateDelete(data, meta)
}

func resourceTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.TemplateCustomizeDiff(data, meta)
}