// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointAPNSChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointAPNSChannelCreate,
		Read:   resourcePinpointAPNSChannelRead,
		Update: resourcePinpointAPNSChannelUpdate,
		Delete: resourcePinpointAPNSChannelDelete,

		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"default_authentication_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"team_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePinpointAPNSChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::APNSChannel", ResourcePinpointAPNSChannel(), data, meta)
}

func resourcePinpointAPNSChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::APNSChannel", ResourcePinpointAPNSChannel(), data, meta)
}

func resourcePinpointAPNSChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::APNSChannel", ResourcePinpointAPNSChannel(), data, meta)
}

func resourcePinpointAPNSChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::APNSChannel", data, meta)
}