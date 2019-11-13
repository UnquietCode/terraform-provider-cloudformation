// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointSmsTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointSmsTemplateExists,
		Read:   resourcePinpointSmsTemplateRead,
		Create: resourcePinpointSmsTemplateCreate,
		Update: resourcePinpointSmsTemplateUpdate,
		Delete: resourcePinpointSmsTemplateDelete,
		CustomizeDiff: resourcePinpointSmsTemplateCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"body": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointSmsTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointSmsTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::SmsTemplate", data, meta)
}

func resourcePinpointSmsTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}