// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointPushTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointPushTemplateCreate,
		Read:   resourcePinpointPushTemplateRead,
		Update: resourcePinpointPushTemplateUpdate,
		Delete: resourcePinpointPushTemplateDelete,

		Schema: map[string]*schema.Schema{
			"gcm": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"baidu": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"template_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"adm": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAndroidPushNotificationTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"apns": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateAPNSPushNotificationTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"default": {
				Type: schema.TypeList,
				Elem: propertyPushTemplateDefaultPushNotificationTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourcePinpointPushTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::PushTemplate", data, meta)
}

func resourcePinpointPushTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::PushTemplate", data, meta)
}

func resourcePinpointPushTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::PushTemplate", data, meta)
}

func resourcePinpointPushTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::PushTemplate", data, meta)
}