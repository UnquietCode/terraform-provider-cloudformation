// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TransitGatewayRouteTableAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TransitGatewayRouteTableAssociationExists,
		Read: resourceEC2TransitGatewayRouteTableAssociationRead,
		Create: resourceEC2TransitGatewayRouteTableAssociationCreate,
		Update: resourceEC2TransitGatewayRouteTableAssociationUpdate,
		Delete: resourceEC2TransitGatewayRouteTableAssociationDelete,
		CustomizeDiff: resourceEC2TransitGatewayRouteTableAssociationCustomizeDiff,
		
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

func resourceEC2TransitGatewayRouteTableAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayRouteTableAssociation", ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayRouteTableAssociation", ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayRouteTableAssociation", ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayRouteTableAssociation", data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
