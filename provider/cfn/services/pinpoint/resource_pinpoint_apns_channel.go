// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointAPNSChannelType string = "AWS::Pinpoint::APNSChannel"

var pinpointAPNSChannelProperties map[string]string = map[string]string{
	"bundle_id": "BundleId",
	"private_key": "PrivateKey",
	"enabled": "Enabled",
	"default_authentication_method": "DefaultAuthenticationMethod",
	"token_key": "TokenKey",
	"application_id": "ApplicationId",
	"team_id": "TeamId",
	"certificate": "Certificate",
	"token_key_id": "TokenKeyId",
}

func ResourcePinpointAPNSChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointAPNSChannelExists,
		Read: resourcePinpointAPNSChannelRead,
		Create: resourcePinpointAPNSChannelCreate,
		Update: resourcePinpointAPNSChannelUpdate,
		Delete: resourcePinpointAPNSChannelDelete,
		CustomizeDiff: resourcePinpointAPNSChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"default_authentication_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"team_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token_key_id": {
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

func resourcePinpointAPNSChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointAPNSChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointAPNSChannelType, ResourcePinpointAPNSChannel(), data, meta)
}

func resourcePinpointAPNSChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointAPNSChannelType, ResourcePinpointAPNSChannel(), data, pinpointAPNSChannelProperties, meta)
}

func resourcePinpointAPNSChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointAPNSChannelType, ResourcePinpointAPNSChannel(), data, pinpointAPNSChannelProperties, meta)
}

func resourcePinpointAPNSChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointAPNSChannelType, data, meta)
}

func resourcePinpointAPNSChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointAPNSChannelType, data, meta)
}
