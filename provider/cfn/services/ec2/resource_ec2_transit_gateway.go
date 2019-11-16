// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const eC2TransitGatewayType string = "AWS::EC2::TransitGateway"

var eC2TransitGatewayProperties map[string]string = map[string]string{
	"default_route_table_propagation": "DefaultRouteTablePropagation",
	"description": "Description",
	"auto_accept_shared_attachments": "AutoAcceptSharedAttachments",
	"default_route_table_association": "DefaultRouteTableAssociation",
	"vpn_ecmp_support": "VpnEcmpSupport",
	"dns_support": "DnsSupport",
	"amazon_side_asn": "AmazonSideAsn",
	"tags": "Tags",
}

func ResourceEC2TransitGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TransitGatewayExists,
		Read: resourceEC2TransitGatewayRead,
		Create: resourceEC2TransitGatewayCreate,
		Update: resourceEC2TransitGatewayUpdate,
		Delete: resourceEC2TransitGatewayDelete,
		CustomizeDiff: resourceEC2TransitGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"default_route_table_propagation": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_accept_shared_attachments": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_route_table_association": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpn_ecmp_support": {
				Type: schema.TypeString,
				Optional: true,
			},
			"dns_support": {
				Type: schema.TypeString,
				Optional: true,
			},
			"amazon_side_asn": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TransitGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TransitGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayType, ResourceEC2TransitGateway(), data, meta)
}

func resourceEC2TransitGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayType, ResourceEC2TransitGateway(), data, eC2TransitGatewayProperties, meta)
}

func resourceEC2TransitGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayType, ResourceEC2TransitGateway(), data, eC2TransitGatewayProperties, meta)
}

func resourceEC2TransitGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayType, data, meta)
}

func resourceEC2TransitGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayType, data, meta)
}
