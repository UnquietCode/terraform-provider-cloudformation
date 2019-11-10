// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEmailTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailTemplateCreate,
		Read:   resourceEmailTemplateRead,
		Update: resourceEmailTemplateUpdate,
		Delete: resourceEmailTemplateDelete,

		Schema: map[string]*schema.Schema{
			"html_part": {
				Type: schema.TypeString,
				Required: false,
			},
			"text_part": {
				Type: schema.TypeString,
				Required: false,
			},
			"template_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subject": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceEmailTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourceEmailTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourceEmailTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourceEmailTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EmailTemplate", data, meta)
}