// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-app.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointAppType string = "AWS::Pinpoint::App"

func ResourcePinpointApp() *schema.Resource {
	return &schema.Resource{
		Read: resourcePinpointAppRead,
		Create: resourcePinpointAppCreate,
		Update: resourcePinpointAppUpdate,
		Delete: resourcePinpointAppDelete,
		CustomizeDiff: resourcePinpointAppCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
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

func resourcePinpointAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointAppType, ResourcePinpointApp(), data, meta)
}

func resourcePinpointAppCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointAppType, ResourcePinpointApp(), data, meta)
}

func resourcePinpointAppUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointAppType, ResourcePinpointApp(), data, meta)
}

func resourcePinpointAppDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointAppType, data, meta)
}

func resourcePinpointAppCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointAppType, data, meta)
}
