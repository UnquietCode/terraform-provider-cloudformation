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

func ResourceEC2VPNConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPNConnectionCreate,
		Read:   resourceEC2VPNConnectionRead,
		Update: resourceEC2VPNConnectionUpdate,
		Delete: resourceEC2VPNConnectionDelete,

		Schema: map[string]*schema.Schema{
			"customer_gateway_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes_only": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"transit_gateway_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpn_tunnel_options_specifications": {
				Type: schema.TypeSet,
				Elem: propertyVpnTunnelOptionsSpecification(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPNConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNConnection", data, meta)
}

func resourceEC2VPNConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNConnection", data, meta)
}

func resourceEC2VPNConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNConnection", data, meta)
}

func resourceEC2VPNConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNConnection", data, meta)
}