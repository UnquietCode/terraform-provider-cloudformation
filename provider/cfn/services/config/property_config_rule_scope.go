// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfigRuleScope() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compliance_resource_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"compliance_resource_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"tag_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_value": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}