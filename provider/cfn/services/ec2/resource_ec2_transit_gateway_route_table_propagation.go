// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetablepropagation.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TransitGatewayRouteTablePropagation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TransitGatewayRouteTablePropagationExists,
		Read: resourceEC2TransitGatewayRouteTablePropagationRead,
		Create: resourceEC2TransitGatewayRouteTablePropagationCreate,
		Update: resourceEC2TransitGatewayRouteTablePropagationUpdate,
		Delete: resourceEC2TransitGatewayRouteTablePropagationDelete,
		CustomizeDiff: resourceEC2TransitGatewayRouteTablePropagationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"transit_gateway_route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"transit_gateway_attachment_id": {
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

func resourceEC2TransitGatewayRouteTablePropagationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayRouteTablePropagation", ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayRouteTablePropagation", ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayRouteTablePropagation", ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayRouteTablePropagation", data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
