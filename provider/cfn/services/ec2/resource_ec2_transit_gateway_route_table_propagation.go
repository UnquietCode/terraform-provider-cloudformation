// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetablepropagation.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayRouteTablePropagationType string = "AWS::EC2::TransitGatewayRouteTablePropagation"

func ResourceEC2TransitGatewayRouteTablePropagation() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2TransitGatewayRouteTablePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayRouteTablePropagationType, ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayRouteTablePropagationType, ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayRouteTablePropagationType, ResourceEC2TransitGatewayRouteTablePropagation(), data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayRouteTablePropagationType, data, meta)
}

func resourceEC2TransitGatewayRouteTablePropagationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayRouteTablePropagationType, data, meta)
}
