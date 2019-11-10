// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransitGatewayRouteTablePropagation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransitGatewayRouteTablePropagationCreate,
		Read:   resourceTransitGatewayRouteTablePropagationRead,
		Update: resourceTransitGatewayRouteTablePropagationUpdate,
		Delete: resourceTransitGatewayRouteTablePropagationDelete,

		Schema: map[string]*schema.Schema{
			"transit_gateway_route_table_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transit_gateway_attachment_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTransitGatewayRouteTablePropagationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayRouteTablePropagation", data, meta)
}

func resourceTransitGatewayRouteTablePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayRouteTablePropagation", data, meta)
}

func resourceTransitGatewayRouteTablePropagationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayRouteTablePropagation", data, meta)
}

func resourceTransitGatewayRouteTablePropagationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayRouteTablePropagation", data, meta)
}