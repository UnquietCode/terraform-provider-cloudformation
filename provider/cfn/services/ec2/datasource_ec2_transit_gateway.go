// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayType string = "AWS::EC2::TransitGateway"

func DatasourceEC2TransitGateway() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2TransitGatewayRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2TransitGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayType, DatasourceEC2TransitGateway(), data, meta)
}
