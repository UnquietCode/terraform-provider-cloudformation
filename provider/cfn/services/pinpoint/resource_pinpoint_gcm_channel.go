// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointGCMChannelType string = "AWS::Pinpoint::GCMChannel"

func ResourcePinpointGCMChannel() *schema.Resource {
	return &schema.Resource{
		Read: resourcePinpointGCMChannelRead,
		Create: resourcePinpointGCMChannelCreate,
		Update: resourcePinpointGCMChannelUpdate,
		Delete: resourcePinpointGCMChannelDelete,
		CustomizeDiff: resourcePinpointGCMChannelCustomizeDiff,
		
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointGCMChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointGCMChannelType, ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointGCMChannelType, ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointGCMChannelType, ResourcePinpointGCMChannel(), data, meta)
}

func resourcePinpointGCMChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointGCMChannelType, data, meta)
}

func resourcePinpointGCMChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointGCMChannelType, data, meta)
}
