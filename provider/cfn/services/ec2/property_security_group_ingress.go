// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var securityGroupIngressProperties map[string]string = map[string]string{
	"cidr_ip": "CidrIp",
	"cidr_ipv6": "CidrIpv6",
	"description": "Description",
	"from_port": "FromPort",
	"ip_protocol": "IpProtocol",
	"source_prefix_list_id": "SourcePrefixListId",
	"source_security_group_id": "SourceSecurityGroupId",
	"source_security_group_name": "SourceSecurityGroupName",
	"source_security_group_owner_id": "SourceSecurityGroupOwnerId",
	"to_port": "ToPort",
}

func propertySecurityGroupIngress(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
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
		},
	}
}
