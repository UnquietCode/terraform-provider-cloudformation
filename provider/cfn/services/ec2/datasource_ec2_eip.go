// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2EIPType string = "AWS::EC2::EIP"

func DatasourceEC2EIP() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2EIPRead,
		
		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"public_ipv4_pool": {
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

func datasourceEC2EIPRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2EIPType, DatasourceEC2EIP(), data, meta)
}
