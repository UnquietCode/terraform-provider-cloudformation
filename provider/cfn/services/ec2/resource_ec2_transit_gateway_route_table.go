// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TransitGatewayRouteTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2TransitGatewayRouteTableCreate,
		Read:   resourceEC2TransitGatewayRouteTableRead,
		Update: resourceEC2TransitGatewayRouteTableUpdate,
		Delete: resourceEC2TransitGatewayRouteTableDelete,

		Schema: map[string]*schema.Schema{
			"transit_gateway_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TransitGatewayRouteTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayRouteTable", data, meta)
}

func resourceEC2TransitGatewayRouteTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayRouteTable", data, meta)
}

func resourceEC2TransitGatewayRouteTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayRouteTable", data, meta)
}

func resourceEC2TransitGatewayRouteTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayRouteTable", data, meta)
}