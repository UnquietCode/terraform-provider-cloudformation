// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html

package codebuild

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var projectSourceProperties map[string]string = map[string]string{
	"type": "Type",
	"report_build_status": "ReportBuildStatus",
	"auth": "Auth",
	"source_identifier": "SourceIdentifier",
	"build_spec": "BuildSpec",
	"git_clone_depth": "GitCloneDepth",
	"git_submodules_config": "GitSubmodulesConfig",
	"insecure_ssl": "InsecureSsl",
	"location": "Location",
}

func propertyProjectSource(extras...string) *schema.Resource {
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
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"report_build_status": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"auth": {
				Type: schema.TypeList,
				Elem: propertyProjectSourceAuth(),
				Optional: true,
				MaxItems: 1,
			},
			"source_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"build_spec": {
				Type: schema.TypeString,
				Optional: true,
			},
			"git_clone_depth": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"git_submodules_config": {
				Type: schema.TypeList,
				Elem: propertyProjectGitSubmodulesConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"insecure_ssl": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"location": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
