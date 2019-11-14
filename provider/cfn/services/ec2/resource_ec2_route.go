// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceEC2RouteExists,
		Read: resourceEC2RouteRead,
		Create: resourceEC2RouteCreate,
		Update: resourceEC2RouteUpdate,
		Delete: resourceEC2RouteDelete,
		CustomizeDiff: resourceEC2RouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_ipv6_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"egress_only_internet_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"nat_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"transit_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_peering_connection_id": {
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

func resourceEC2RouteExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Route", ResourceEC2Route(), data, meta)
}

func resourceEC2RouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Route", ResourceEC2Route(), data, meta)
}

func resourceEC2RouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Route", ResourceEC2Route(), data, meta)
}

func resourceEC2RouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Route", data, meta)
}

func resourceEC2RouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
