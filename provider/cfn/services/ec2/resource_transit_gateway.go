// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransitGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransitGatewayCreate,
		Read:   resourceTransitGatewayRead,
		Update: resourceTransitGatewayUpdate,
		Delete: resourceTransitGatewayDelete,

		Schema: map[string]*schema.Schema{
			"default_route_table_propagation": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"auto_accept_shared_attachments": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"default_route_table_association": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpn_ecmp_support": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"dns_support": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"amazon_side_asn": {
				Type: schema.TypeInt,
				Required: false,
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

func resourceTransitGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGateway", data, meta)
}

func resourceTransitGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGateway", data, meta)
}

func resourceTransitGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGateway", data, meta)
}

func resourceTransitGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGateway", data, meta)
}