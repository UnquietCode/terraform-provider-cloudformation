// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnsvoipsandboxchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointAPNSVoipSandboxChannelType string = "AWS::Pinpoint::APNSVoipSandboxChannel"

var pinpointAPNSVoipSandboxChannelProperties map[string]string = map[string]string{
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

func ResourcePinpointAPNSVoipSandboxChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointAPNSVoipSandboxChannelExists,
		Read: resourcePinpointAPNSVoipSandboxChannelRead,
		Create: resourcePinpointAPNSVoipSandboxChannelCreate,
		Update: resourcePinpointAPNSVoipSandboxChannelUpdate,
		Delete: resourcePinpointAPNSVoipSandboxChannelDelete,
		CustomizeDiff: resourcePinpointAPNSVoipSandboxChannelCustomizeDiff,
		
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

func resourcePinpointAPNSVoipSandboxChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointAPNSVoipSandboxChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointAPNSVoipSandboxChannelType, ResourcePinpointAPNSVoipSandboxChannel(), data, meta)
}

func resourcePinpointAPNSVoipSandboxChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointAPNSVoipSandboxChannelType, ResourcePinpointAPNSVoipSandboxChannel(), data, pinpointAPNSVoipSandboxChannelProperties, meta)
}

func resourcePinpointAPNSVoipSandboxChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointAPNSVoipSandboxChannelType, ResourcePinpointAPNSVoipSandboxChannel(), data, pinpointAPNSVoipSandboxChannelProperties, meta)
}

func resourcePinpointAPNSVoipSandboxChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointAPNSVoipSandboxChannelType, data, meta)
}

func resourcePinpointAPNSVoipSandboxChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointAPNSVoipSandboxChannelType, data, meta)
}
