package misc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func DatasourceTemplateData() *schema.Resource {
	return &schema.Resource{
		Read: 	resourceTemplateRead,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
      "resources": {
        Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
        Required: true,
      },
			"json": {
				Type: schema.TypeString,
				Computed: true,
        Sensitive: true,
			},
		},
	}
}

func resourceTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.TemplateRead(data, meta)
}