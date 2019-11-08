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

func ResourcePinpointSMSChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointSMSChannelCreate,
		Read:   resourcePinpointSMSChannelRead,
		Update: resourcePinpointSMSChannelUpdate,
		Delete: resourcePinpointSMSChannelDelete,

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

func resourcePinpointSMSChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourcePinpointSMSChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourcePinpointSMSChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::SMSChannel", data, meta)
}

func resourcePinpointSMSChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::SMSChannel", data, meta)
}