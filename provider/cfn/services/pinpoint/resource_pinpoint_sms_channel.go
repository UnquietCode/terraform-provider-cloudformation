// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointSMSChannelType string = "AWS::Pinpoint::SMSChannel"

func ResourcePinpointSMSChannel() *schema.Resource {
	return &schema.Resource{
		Read: resourcePinpointSMSChannelRead,
		Create: resourcePinpointSMSChannelCreate,
		Update: resourcePinpointSMSChannelUpdate,
		Delete: resourcePinpointSMSChannelDelete,
		CustomizeDiff: resourcePinpointSMSChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"short_code": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"sender_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointSMSChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointSMSChannelType, ResourcePinpointSMSChannel(), data, meta)
}

func resourcePinpointSMSChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointSMSChannelType, ResourcePinpointSMSChannel(), data, meta)
}

func resourcePinpointSMSChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointSMSChannelType, ResourcePinpointSMSChannel(), data, meta)
}

func resourcePinpointSMSChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointSMSChannelType, data, meta)
}

func resourcePinpointSMSChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointSMSChannelType, data, meta)
}
