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

func ResourcePinpointAPNSVoipChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointAPNSVoipChannelCreate,
		Read:   resourcePinpointAPNSVoipChannelRead,
		Update: resourcePinpointAPNSVoipChannelUpdate,
		Delete: resourcePinpointAPNSVoipChannelDelete,

		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"private_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"default_authentication_method": {
				Type: schema.TypeString,
				Required: false,
			},
			"token_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"team_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"certificate": {
				Type: schema.TypeString,
				Required: false,
			},
			"token_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourcePinpointAPNSVoipChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::APNSVoipChannel", data, meta)
}

func resourcePinpointAPNSVoipChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::APNSVoipChannel", data, meta)
}

func resourcePinpointAPNSVoipChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::APNSVoipChannel", data, meta)
}

func resourcePinpointAPNSVoipChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::APNSVoipChannel", data, meta)
}