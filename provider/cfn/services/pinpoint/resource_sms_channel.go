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

func ResourceSMSChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSMSChannelCreate,
		Read:   resourceSMSChannelRead,
		Update: resourceSMSChannelUpdate,
		Delete: resourceSMSChannelDelete,

		Schema: map[string]*schema.Schema{
			"short_code": {
				Type: schema.TypeString,
				Required: false,
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
			"sender_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceSMSChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourceSMSChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourceSMSChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourceSMSChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::SMSChannel", data, meta)
}