// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointGCMChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointGCMChannelCreate,
		Read:   resourcePinpointGCMChannelRead,
		Update: resourcePinpointGCMChannelUpdate,
		Delete: resourcePinpointGCMChannelDelete,

		Schema: map[string]*schema.Schema{
			"api_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointGCMChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::GCMChannel", ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::GCMChannel", ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::GCMChannel", ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::GCMChannel", data, meta)
}