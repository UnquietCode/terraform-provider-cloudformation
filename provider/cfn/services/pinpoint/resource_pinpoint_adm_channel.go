// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointADMChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointADMChannelCreate,
		Read:   resourcePinpointADMChannelRead,
		Update: resourcePinpointADMChannelUpdate,
		Delete: resourcePinpointADMChannelDelete,

		Schema: map[string]*schema.Schema{
			"client_secret": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"client_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointADMChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::ADMChannel", ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::ADMChannel", ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::ADMChannel", ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::ADMChannel", data, meta)
}