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

func ResourceVPNGatewayRoutePropagation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPNGatewayRoutePropagationCreate,
		Read:   resourceVPNGatewayRoutePropagationRead,
		Update: resourceVPNGatewayRoutePropagationUpdate,
		Delete: resourceVPNGatewayRoutePropagationDelete,

		Schema: map[string]*schema.Schema{
			"route_table_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVPNGatewayRoutePropagationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceVPNGatewayRoutePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceVPNGatewayRoutePropagationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceVPNGatewayRoutePropagationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}