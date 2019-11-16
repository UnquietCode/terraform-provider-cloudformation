// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-autobranchcreationconfig.html

package amplify

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var appAutoBranchCreationConfigProperties map[string]string = map[string]string{
	"environment_variables": "EnvironmentVariables",
	"enable_auto_branch_creation": "EnableAutoBranchCreation",
	"pull_request_environment_name": "PullRequestEnvironmentName",
	"auto_branch_creation_patterns": "AutoBranchCreationPatterns",
	"enable_pull_request_preview": "EnablePullRequestPreview",
	"enable_auto_build": "EnableAutoBuild",
	"build_spec": "BuildSpec",
	"stage": "Stage",
	"basic_auth_config": "BasicAuthConfig",
}

func propertyAppAutoBranchCreationConfig(extras...string) *schema.Resource {
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
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Optional: true,
			},
			"enable_auto_branch_creation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"pull_request_environment_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_branch_creation_patterns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enable_pull_request_preview": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enable_auto_build": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"build_spec": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stage": {
				Type: schema.TypeString,
				Optional: true,
			},
			"basic_auth_config": {
				Type: schema.TypeSet,
				Elem: propertyAppBasicAuthConfig(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
