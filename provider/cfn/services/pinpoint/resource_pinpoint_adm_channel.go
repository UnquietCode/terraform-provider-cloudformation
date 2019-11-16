// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointADMChannelType string = "AWS::Pinpoint::ADMChannel"

func ResourcePinpointADMChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointADMChannelExists,
		Read: resourcePinpointADMChannelRead,
		Create: resourcePinpointADMChannelCreate,
		Update: resourcePinpointADMChannelUpdate,
		Delete: resourcePinpointADMChannelDelete,
		CustomizeDiff: resourcePinpointADMChannelCustomizeDiff,
		
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointADMChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointADMChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointADMChannelType, ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointADMChannelType, ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointADMChannelType, ResourcePinpointADMChannel(), data, meta)
}

func resourcePinpointADMChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointADMChannelType, data, meta)
}

func resourcePinpointADMChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointADMChannelType, data, meta)
}
