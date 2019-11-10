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

func ResourcePinpointEmailTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointEmailTemplateCreate,
		Read:   resourcePinpointEmailTemplateRead,
		Update: resourcePinpointEmailTemplateUpdate,
		Delete: resourcePinpointEmailTemplateDelete,

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

func resourcePinpointEmailTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourcePinpointEmailTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourcePinpointEmailTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EmailTemplate", data, meta)
}

func resourcePinpointEmailTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EmailTemplate", data, meta)
}