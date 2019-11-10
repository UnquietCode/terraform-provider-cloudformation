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

func ResourceTransitGatewayRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransitGatewayRouteCreate,
		Read:   resourceTransitGatewayRouteRead,
		Update: resourceTransitGatewayRouteUpdate,
		Delete: resourceTransitGatewayRouteDelete,

		Schema: map[string]*schema.Schema{
			"transit_gateway_route_table_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"blackhole": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"transit_gateway_attachment_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceTransitGatewayRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayRoute", data, meta)
}

func resourceTransitGatewayRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayRoute", data, meta)
}

func resourceTransitGatewayRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayRoute", data, meta)
}

func resourceTransitGatewayRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayRoute", data, meta)
}