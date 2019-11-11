// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TransitGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2TransitGatewayCreate,
		Read:   resourceEC2TransitGatewayRead,
		Delete: resourceEC2TransitGatewayDelete,

		Schema: map[string]*schema.Schema{
			"default_route_table_propagation": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"auto_accept_shared_attachments": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_route_table_association": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpn_ecmp_support": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dns_support": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"amazon_side_asn": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TransitGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGateway", ResourceEC2TransitGateway(), data, meta)
}

func resourceEC2TransitGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGateway", ResourceEC2TransitGateway(), data, meta)
}

func resourceEC2TransitGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGateway", ResourceEC2TransitGateway(), data, meta)
}

func resourceEC2TransitGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGateway", data, meta)
}