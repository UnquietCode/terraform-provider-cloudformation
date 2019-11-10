// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySecurityGroupEgress(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Required: false,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"destination_prefix_list_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"destination_security_group_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"from_port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"to_port": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}