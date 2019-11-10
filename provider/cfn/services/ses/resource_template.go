// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCreate,
		Read:   resourceTemplateRead,
		Update: resourceTemplateUpdate,
		Delete: resourceTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template": {
				Type: schema.TypeList,
				Elem: propertyTemplateTemplate(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::Template", data, meta)
}

func resourceTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::Template", data, meta)
}

func resourceTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::Template", data, meta)
}

func resourceTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::Template", data, meta)
}