// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2CustomerGatewayType string = "AWS::EC2::CustomerGateway"

func DatasourceEC2CustomerGateway() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2CustomerGatewayRead,
		
		Schema: map[string]*schema.Schema{
			"bgp_asn": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ip_address": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"type": {
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

func datasourceEC2CustomerGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2CustomerGatewayType, DatasourceEC2CustomerGateway(), data, meta)
}
