// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package amplify

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAppAutoBranchCreationConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Required: false,
			},
			"enable_auto_branch_creation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"pull_request_environment_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"auto_branch_creation_patterns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"enable_pull_request_preview": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enable_auto_build": {
				Type: schema.TypeBool,
				Required: false,
			},
			"build_spec": {
				Type: schema.TypeString,
				Required: false,
			},
			"stage": {
				Type: schema.TypeString,
				Required: false,
			},
			"basic_auth_config": {
				Type: schema.TypeList,
				Elem: propertyAppBasicAuthConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}