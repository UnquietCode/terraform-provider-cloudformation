// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyIngress() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Required: false,
			},
			"ec2_security_group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ec2_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}