// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourcePinpointSmsTemplateCreate,
		Read:   resourcePinpointSmsTemplateRead,
		Update: resourcePinpointSmsTemplateUpdate,
		Delete: resourcePinpointSmsTemplateDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				Optional: true,
			},
		},
	}
}

func resourcePinpointSmsTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::SmsTemplate", ResourcePinpointSmsTemplate(), data, meta)
}

func resourcePinpointSmsTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::SmsTemplate", data, meta)
}