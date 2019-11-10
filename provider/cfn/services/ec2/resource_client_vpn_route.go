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

func ResourceClientVpnRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientVpnRouteCreate,
		Read:   resourceClientVpnRouteRead,
		Update: resourceClientVpnRouteUpdate,
		Delete: resourceClientVpnRouteDelete,

		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_vpc_subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceClientVpnRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnRoute", data, meta)
}

func resourceClientVpnRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnRoute", data, meta)
}

func resourceClientVpnRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnRoute", data, meta)
}

func resourceClientVpnRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnRoute", data, meta)
}