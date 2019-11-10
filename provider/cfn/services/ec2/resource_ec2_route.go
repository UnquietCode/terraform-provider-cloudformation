// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2Route() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2RouteCreate,
		Read:   resourceEC2RouteRead,
		Update: resourceEC2RouteUpdate,
		Delete: resourceEC2RouteDelete,

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

func resourceEC2RouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Route", data, meta)
}

func resourceEC2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Route", data, meta)
}

func resourceEC2RouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Route", data, meta)
}

func resourceEC2RouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Route", data, meta)
}