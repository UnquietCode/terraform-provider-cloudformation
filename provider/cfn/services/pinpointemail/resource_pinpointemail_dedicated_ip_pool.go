// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-dedicatedippool.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailDedicatedIpPoolType string = "AWS::PinpointEmail::DedicatedIpPool"

func ResourcePinpointEmailDedicatedIpPool() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEmailDedicatedIpPoolExists,
		Read: resourcePinpointEmailDedicatedIpPoolRead,
		Create: resourcePinpointEmailDedicatedIpPoolCreate,
		Update: resourcePinpointEmailDedicatedIpPoolUpdate,
		Delete: resourcePinpointEmailDedicatedIpPoolDelete,
		CustomizeDiff: resourcePinpointEmailDedicatedIpPoolCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyDedicatedIpPoolTags(),
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

func resourcePinpointEmailDedicatedIpPoolExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailDedicatedIpPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailDedicatedIpPoolType, ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointEmailDedicatedIpPoolType, ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointEmailDedicatedIpPoolType, ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointEmailDedicatedIpPoolType, data, meta)
}

func resourcePinpointEmailDedicatedIpPoolCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointEmailDedicatedIpPoolType, data, meta)
}
