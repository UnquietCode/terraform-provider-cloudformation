// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceADMChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceADMChannelCreate,
		Read:   resourceADMChannelRead,
		Update: resourceADMChannelUpdate,
		Delete: resourceADMChannelDelete,

		Schema: map[string]*schema.Schema{
			"client_secret": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
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
		},
	}
}

func resourceADMChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::ADMChannel", data, meta)
}

func resourceADMChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::ADMChannel", data, meta)
}

func resourceADMChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::ADMChannel", data, meta)
}

func resourceADMChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::ADMChannel", data, meta)
}