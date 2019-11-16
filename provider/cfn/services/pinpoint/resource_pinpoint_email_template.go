// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailTemplateType string = "AWS::Pinpoint::EmailTemplate"

var pinpointEmailTemplateProperties map[string]string = map[string]string{
	"html_part": "HtmlPart",
	"text_part": "TextPart",
	"template_name": "TemplateName",
	"subject": "Subject",
	"tags": "Tags",
}

func ResourcePinpointEmailTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEmailTemplateExists,
		Read: resourcePinpointEmailTemplateRead,
		Create: resourcePinpointEmailTemplateCreate,
		Update: resourcePinpointEmailTemplateUpdate,
		Delete: resourcePinpointEmailTemplateDelete,
		CustomizeDiff: resourcePinpointEmailTemplateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"html_part": {
				Type: schema.TypeString,
				Optional: true,
			},
			"text_part": {
				Type: schema.TypeString,
				Optional: true,
			},
			"template_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"subject": {
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

func resourcePinpointEmailTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailTemplateType, ResourcePinpointEmailTemplate(), data, meta)
}

func resourcePinpointEmailTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointEmailTemplateType, ResourcePinpointEmailTemplate(), data, pinpointEmailTemplateProperties, meta)
}

func resourcePinpointEmailTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointEmailTemplateType, ResourcePinpointEmailTemplate(), data, pinpointEmailTemplateProperties, meta)
}

func resourcePinpointEmailTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointEmailTemplateType, data, meta)
}

func resourcePinpointEmailTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointEmailTemplateType, data, meta)
}
