// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayRouteTableAssociationType string = "AWS::EC2::TransitGatewayRouteTableAssociation"

func ResourceEC2TransitGatewayRouteTableAssociation() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2TransitGatewayRouteTableAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayRouteTableAssociationType, ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayRouteTableAssociationType, ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayRouteTableAssociationType, ResourceEC2TransitGatewayRouteTableAssociation(), data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayRouteTableAssociationType, data, meta)
}

func resourceEC2TransitGatewayRouteTableAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayRouteTableAssociationType, data, meta)
}
