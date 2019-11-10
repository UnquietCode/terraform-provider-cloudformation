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

func ResourceEmailChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailChannelCreate,
		Read:   resourceEmailChannelRead,
		Update: resourceEmailChannelUpdate,
		Delete: resourceEmailChannelDelete,

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

func resourceEmailChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourceEmailChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourceEmailChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::EmailChannel", data, meta)
}

func resourceEmailChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::EmailChannel", data, meta)
}