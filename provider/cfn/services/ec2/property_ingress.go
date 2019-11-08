// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyIngress() *schema.Resource {
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
			"from_port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_prefix_list_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_security_group_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_security_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"to_port": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}