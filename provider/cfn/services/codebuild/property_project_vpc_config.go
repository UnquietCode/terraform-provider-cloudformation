// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProjectVpcConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnets": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}