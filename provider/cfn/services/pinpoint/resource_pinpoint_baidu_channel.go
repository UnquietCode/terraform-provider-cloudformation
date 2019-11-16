// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointBaiduChannelType string = "AWS::Pinpoint::BaiduChannel"

var pinpointBaiduChannelProperties map[string]string = map[string]string{
	"secret_key": "SecretKey",
	"api_key": "ApiKey",
	"enabled": "Enabled",
	"application_id": "ApplicationId",
}

func ResourcePinpointBaiduChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointBaiduChannelExists,
		Read: resourcePinpointBaiduChannelRead,
		Create: resourcePinpointBaiduChannelCreate,
		Update: resourcePinpointBaiduChannelUpdate,
		Delete: resourcePinpointBaiduChannelDelete,
		CustomizeDiff: resourcePinpointBaiduChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"secret_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_key": {
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointBaiduChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointBaiduChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointBaiduChannelType, ResourcePinpointBaiduChannel(), data, meta)
}

func resourcePinpointBaiduChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointBaiduChannelType, ResourcePinpointBaiduChannel(), data, pinpointBaiduChannelProperties, meta)
}

func resourcePinpointBaiduChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointBaiduChannelType, ResourcePinpointBaiduChannel(), data, pinpointBaiduChannelProperties, meta)
}

func resourcePinpointBaiduChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointBaiduChannelType, data, meta)
}

func resourcePinpointBaiduChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointBaiduChannelType, data, meta)
}
