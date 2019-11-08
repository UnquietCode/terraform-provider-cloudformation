// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyOrganizationManagedRuleMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_key_scope": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_value_scope": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"resource_id_scope": {
				Type: schema.TypeString,
				Required: false,
			},
			"rule_identifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_types_scope": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_parameters": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}