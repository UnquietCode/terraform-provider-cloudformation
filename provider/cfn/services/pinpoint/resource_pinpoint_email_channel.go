// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailChannelType string = "AWS::Pinpoint::EmailChannel"

var pinpointEmailChannelProperties map[string]string = map[string]string{
	"configuration_set": "ConfigurationSet",
	"from_address": "FromAddress",
	"enabled": "Enabled",
	"application_id": "ApplicationId",
	"identity": "Identity",
	"role_arn": "RoleArn",
}

func ResourcePinpointEmailChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEmailChannelExists,
		Read: resourcePinpointEmailChannelRead,
		Create: resourcePinpointEmailChannelCreate,
		Update: resourcePinpointEmailChannelUpdate,
		Delete: resourcePinpointEmailChannelDelete,
		CustomizeDiff: resourcePinpointEmailChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"configuration_set": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_address": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"identity": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
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

func resourcePinpointEmailChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailChannelType, ResourcePinpointEmailChannel(), data, meta)
}

func resourcePinpointEmailChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointEmailChannelType, ResourcePinpointEmailChannel(), data, pinpointEmailChannelProperties, meta)
}

func resourcePinpointEmailChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointEmailChannelType, ResourcePinpointEmailChannel(), data, pinpointEmailChannelProperties, meta)
}

func resourcePinpointEmailChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointEmailChannelType, data, meta)
}

func resourcePinpointEmailChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointEmailChannelType, data, meta)
}
