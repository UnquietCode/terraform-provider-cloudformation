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

func ResourceVPNConnectionRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPNConnectionRouteCreate,
		Read:   resourceVPNConnectionRouteRead,
		Update: resourceVPNConnectionRouteUpdate,
		Delete: resourceVPNConnectionRouteDelete,

		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpn_connection_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPNConnectionRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNConnectionRoute", data, meta)
}

func resourceVPNConnectionRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNConnectionRoute", data, meta)
}

func resourceVPNConnectionRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNConnectionRoute", data, meta)
}

func resourceVPNConnectionRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNConnectionRoute", data, meta)
}