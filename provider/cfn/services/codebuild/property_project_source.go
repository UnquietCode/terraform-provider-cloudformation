// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProjectSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"report_build_status": {
				Type: schema.TypeBool,
				Required: false,
			},
			"auth": {
				Type: schema.TypeList,
				Elem: propertyProjectSourceAuth(),
				Required: false,
				MaxItems: 1,
			},
			"source_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"build_spec": {
				Type: schema.TypeString,
				Required: false,
			},
			"git_clone_depth": {
				Type: schema.TypeInt,
				Required: false,
			},
			"git_submodules_config": {
				Type: schema.TypeList,
				Elem: propertyProjectGitSubmodulesConfig(),
				Required: false,
				MaxItems: 1,
			},
			"insecure_ssl": {
				Type: schema.TypeBool,
				Required: false,
			},
			"location": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}