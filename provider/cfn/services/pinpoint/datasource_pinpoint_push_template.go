// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html

package pinpoint

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointPushTemplateType string = "AWS::Pinpoint::PushTemplate"

func DatasourcePinpointPushTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourcePinpointPushTemplateRead,
		
		Schema: map[string]*schema.Schema{
			"gcm": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"baidu": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"template_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"adm": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"apns": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAPNSPushNotificationTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"template_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_substitutions": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateDefaultPushNotificationTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourcePinpointPushTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointPushTemplateType, DatasourcePinpointPushTemplate(), data, meta)
}
