// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCType string = "AWS::EC2::VPC"

func DatasourceEC2VPC() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2VPCRead,
		
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"enable_dns_hostnames": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enable_dns_support": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"instance_tenancy": {
				Type: schema.TypeString,
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

func datasourceEC2VPCRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCType, DatasourceEC2VPC(), data, meta)
}
