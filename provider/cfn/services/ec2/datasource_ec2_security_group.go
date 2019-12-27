// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SecurityGroupType string = "AWS::EC2::SecurityGroup"

func DatasourceEC2SecurityGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2SecurityGroupRead,
		
		Schema: map[string]*schema.Schema{
			"group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_egress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupEgress(),
				Optional: true,
			},
			"security_group_ingress": {
				Type: schema.TypeList,
				Elem: propertySecurityGroupIngress(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"vpc_id": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceEC2SecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SecurityGroupType, DatasourceEC2SecurityGroup(), data, meta)
}
