// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPNGatewayRoutePropagationType string = "AWS::EC2::VPNGatewayRoutePropagation"

func DatasourceEC2VPNGatewayRoutePropagation() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2VPNGatewayRoutePropagationRead,
		
		Schema: map[string]*schema.Schema{
			"route_table_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Required: true,
			},
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

func datasourceEC2VPNGatewayRoutePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNGatewayRoutePropagationType, DatasourceEC2VPNGatewayRoutePropagation(), data, meta)
}
