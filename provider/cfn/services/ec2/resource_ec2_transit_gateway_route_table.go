// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayRouteTableType string = "AWS::EC2::TransitGatewayRouteTable"

func ResourceEC2TransitGatewayRouteTable() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2TransitGatewayRouteTableRead,
		Create: resourceEC2TransitGatewayRouteTableCreate,
		Update: resourceEC2TransitGatewayRouteTableUpdate,
		Delete: resourceEC2TransitGatewayRouteTableDelete,
		CustomizeDiff: resourceEC2TransitGatewayRouteTableCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"transit_gateway_id": {
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

func resourceEC2TransitGatewayRouteTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayRouteTableType, ResourceEC2TransitGatewayRouteTable(), data, meta)
}

func resourceEC2TransitGatewayRouteTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayRouteTableType, ResourceEC2TransitGatewayRouteTable(), data, meta)
}

func resourceEC2TransitGatewayRouteTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayRouteTableType, ResourceEC2TransitGatewayRouteTable(), data, meta)
}

func resourceEC2TransitGatewayRouteTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayRouteTableType, data, meta)
}

func resourceEC2TransitGatewayRouteTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayRouteTableType, data, meta)
}
