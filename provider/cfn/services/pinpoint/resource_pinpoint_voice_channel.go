// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-voicechannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointVoiceChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointVoiceChannelExists,
		Read:   resourcePinpointVoiceChannelRead,
		Create: resourcePinpointVoiceChannelCreate,
		Update: resourcePinpointVoiceChannelUpdate,
		Delete: resourcePinpointVoiceChannelDelete,
		
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
	return plugin.ResourceRead("AWS::Pinpoint::VoiceChannel", ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::VoiceChannel", ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::VoiceChannel", ResourcePinpointVoiceChannel(), data, meta)
}

func resourcePinpointVoiceChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::VoiceChannel", data, meta)
}