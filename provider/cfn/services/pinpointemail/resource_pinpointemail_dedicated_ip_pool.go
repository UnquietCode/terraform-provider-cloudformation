// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-dedicatedippool.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
	return plugin.ResourceRead("AWS::PinpointEmail::DedicatedIpPool", ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::DedicatedIpPool", ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::DedicatedIpPool", ResourcePinpointEmailDedicatedIpPool(), data, meta)
}

func resourcePinpointEmailDedicatedIpPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}

func resourcePinpointEmailDedicatedIpPoolCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::PinpointEmail::DedicatedIpPool", data, meta)
}

