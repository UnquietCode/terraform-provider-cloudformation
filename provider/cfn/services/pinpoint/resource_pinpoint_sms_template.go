// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointSmsTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointSmsTemplateCreate,
		Read:   resourcePinpointSmsTemplateRead,
		Update: resourcePinpointSmsTemplateUpdate,
		Delete: resourcePinpointSmsTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body": {
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

func resourcePinpointSmsTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::SmsTemplate", data, meta)
}

func resourcePinpointSmsTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::SmsTemplate", data, meta)
}

func resourcePinpointSmsTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::SmsTemplate", data, meta)
}

func resourcePinpointSmsTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::SmsTemplate", data, meta)
}