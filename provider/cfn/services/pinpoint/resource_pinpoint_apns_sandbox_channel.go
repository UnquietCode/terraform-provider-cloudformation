// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html

package pinpoint

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointAPNSSandboxChannelType string = "AWS::Pinpoint::APNSSandboxChannel"

func ResourcePinpointAPNSSandboxChannel() *schema.Resource {
	return &schema.Resource{
		Read: resourcePinpointAPNSSandboxChannelRead,
		Create: resourcePinpointAPNSSandboxChannelCreate,
		Update: resourcePinpointAPNSSandboxChannelUpdate,
		Delete: resourcePinpointAPNSSandboxChannelDelete,
		CustomizeDiff: resourcePinpointAPNSSandboxChannelCustomizeDiff,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourcePinpointAPNSSandboxChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointAPNSSandboxChannelType, ResourcePinpointAPNSSandboxChannel(), data, meta)
}

func resourcePinpointAPNSSandboxChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointAPNSSandboxChannelType, ResourcePinpointAPNSSandboxChannel(), data, meta)
}

func resourcePinpointAPNSSandboxChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointAPNSSandboxChannelType, ResourcePinpointAPNSSandboxChannel(), data, meta)
}

func resourcePinpointAPNSSandboxChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointAPNSSandboxChannelType, data, meta)
}

func resourcePinpointAPNSSandboxChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointAPNSSandboxChannelType, data, meta)
}
