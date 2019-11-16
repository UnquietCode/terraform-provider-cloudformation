// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayRouteType string = "AWS::EC2::TransitGatewayRoute"

func ResourceEC2TransitGatewayRoute() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TransitGatewayRouteExists,
		Read: resourceEC2TransitGatewayRouteRead,
		Create: resourceEC2TransitGatewayRouteCreate,
		Update: resourceEC2TransitGatewayRouteUpdate,
		Delete: resourceEC2TransitGatewayRouteDelete,
		CustomizeDiff: resourceEC2TransitGatewayRouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"transit_gateway_route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"blackhole": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"transit_gateway_attachment_id": {
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

func resourceEC2TransitGatewayRouteExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TransitGatewayRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayRouteType, ResourceEC2TransitGatewayRoute(), data, meta)
}

func resourceEC2TransitGatewayRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayRouteType, ResourceEC2TransitGatewayRoute(), data, meta)
}

func resourceEC2TransitGatewayRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayRouteType, ResourceEC2TransitGatewayRoute(), data, meta)
}

func resourceEC2TransitGatewayRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayRouteType, data, meta)
}

func resourceEC2TransitGatewayRouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayRouteType, data, meta)
}
