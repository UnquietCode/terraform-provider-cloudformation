package misc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTemplateData() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCreate,
		Read:   resourceTemplateRead,
		// Update:   resourceTemplateUpdate,
		Delete: resourceTemplateDelete,		
		CustomizeDiff: plugin.TemplateCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"output": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateCreate(data, meta)
}

func resourceTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateRead(data, meta)
}

func resourceTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateUpdate(data, meta)
}

func resourceTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateDelete(data, meta)
}
