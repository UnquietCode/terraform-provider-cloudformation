// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointVoiceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointVoiceChannelCreate,
		Read:   resourcePinpointVoiceChannelRead,
		Update: resourcePinpointVoiceChannelUpdate,
		Delete: resourcePinpointVoiceChannelDelete,

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointVoiceChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::VoiceChannel", data, meta)
}

func resourcePinpointVoiceChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::VoiceChannel", data, meta)
}

func resourcePinpointVoiceChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::VoiceChannel", data, meta)
}

func resourcePinpointVoiceChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::VoiceChannel", data, meta)
}