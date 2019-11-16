// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NatGatewayType string = "AWS::EC2::NatGateway"

var eC2NatGatewayProperties map[string]string = map[string]string{
	"allocation_id": "AllocationId",
	"subnet_id": "SubnetId",
	"tags": "Tags",
}

func ResourceEC2NatGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2NatGatewayExists,
		Read: resourceEC2NatGatewayRead,
		Create: resourceEC2NatGatewayCreate,
		Update: resourceEC2NatGatewayUpdate,
		Delete: resourceEC2NatGatewayDelete,
		CustomizeDiff: resourceEC2NatGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NatGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2NatGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NatGatewayType, ResourceEC2NatGateway(), data, meta)
}

func resourceEC2NatGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NatGatewayType, ResourceEC2NatGateway(), data, eC2NatGatewayProperties, meta)
}

func resourceEC2NatGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NatGatewayType, ResourceEC2NatGateway(), data, eC2NatGatewayProperties, meta)
}

func resourceEC2NatGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NatGatewayType, data, meta)
}

func resourceEC2NatGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NatGatewayType, data, meta)
}
