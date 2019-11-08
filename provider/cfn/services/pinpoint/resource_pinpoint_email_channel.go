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

func ResourcePinpointEmailChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointEmailChannelCreate,
		Read:   resourcePinpointEmailChannelRead,
		Update: resourcePinpointEmailChannelUpdate,
		Delete: resourcePinpointEmailChannelDelete,

		Schema: map[string]*schema.Schema{
			"configuration_set": {
				Type: schema.TypeString,
				Required: false,
			},
			"from_address": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identity": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourcePinpointEmailChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourcePinpointEmailChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourcePinpointEmailChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourcePinpointEmailChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EmailChannel", data, meta)
}