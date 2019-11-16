// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-voicechannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointVoiceChannelType string = "AWS::Pinpoint::VoiceChannel"

func ResourcePinpointVoiceChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointVoiceChannelExists,
		Read: resourcePinpointVoiceChannelRead,
		Create: resourcePinpointVoiceChannelCreate,
		Update: resourcePinpointVoiceChannelUpdate,
		Delete: resourcePinpointVoiceChannelDelete,
		CustomizeDiff: resourcePinpointVoiceChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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

func resourcePinpointVoiceChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointVoiceChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointVoiceChannelType, ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointVoiceChannelType, ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointVoiceChannelType, ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointVoiceChannelType, data, meta)
}

func resourcePinpointVoiceChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointVoiceChannelType, data, meta)
}
