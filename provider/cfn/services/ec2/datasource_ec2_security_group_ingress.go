// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SecurityGroupIngressType string = "AWS::EC2::SecurityGroupIngress"

func DatasourceEC2SecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2SecurityGroupIngressRead,
		
		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_prefix_list_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_owner_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"to_port": {
				Type: schema.TypeInt,
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

func datasourceEC2SecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SecurityGroupIngressType, DatasourceEC2SecurityGroupIngress(), data, meta)
}
