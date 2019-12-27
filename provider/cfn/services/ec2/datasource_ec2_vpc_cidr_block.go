// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCCidrBlockType string = "AWS::EC2::VPCCidrBlock"

func DatasourceEC2VPCCidrBlock() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2VPCCidrBlockRead,
		
		Schema: map[string]*schema.Schema{
			"amazon_provided_ipv6_cidr_block": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
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

func datasourceEC2VPCCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCCidrBlockType, DatasourceEC2VPCCidrBlock(), data, meta)
}
