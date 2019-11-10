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

func ResourceRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouteCreate,
		Read:   resourceRouteRead,
		Update: resourceRouteUpdate,
		Delete: resourceRouteDelete,

		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"destination_ipv6_cidr_block": {
				Type: schema.TypeString,
				Required: false,
			},
			"egress_only_internet_gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"nat_gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"route_table_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transit_gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"vpc_peering_connection_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Route", data, meta)
}

func resourceRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Route", data, meta)
}

func resourceRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Route", data, meta)
}

func resourceRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Route", data, meta)
}