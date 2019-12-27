// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html

package codebuild

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeBuildProjectType string = "AWS::CodeBuild::Project"

func DatasourceCodeBuildProject() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodeBuildProjectRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyProjectVpcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"secondary_sources": {
				Type: schema.TypeList,
				Elem: propertyProjectSource(),
				Optional: true,
			},
			"encryption_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"triggers": {
				Type: schema.TypeList,
				Elem: propertyProjectProjectTriggers(),
				Optional: true,
				MaxItems: 1,
			},
			"secondary_artifacts": {
				Type: schema.TypeList,
				Elem: propertyProjectArtifacts(),
				Optional: true,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertyProjectSource(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"artifacts": {
				Type: schema.TypeList,
				Elem: propertyProjectArtifacts(),
				Required: true,
				MaxItems: 1,
			},
			"badge_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logs_config": {
				Type: schema.TypeList,
				Elem: propertyProjectLogsConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"queued_timeout_in_minutes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyProjectEnvironment(),
				Required: true,
				MaxItems: 1,
			},
			"secondary_source_versions": {
				Type: schema.TypeList,
				Elem: propertyProjectProjectSourceVersion(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"timeout_in_minutes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"cache": {
				Type: schema.TypeList,
				Elem: propertyProjectProjectCache(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodeBuildProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeBuildProjectType, DatasourceCodeBuildProject(), data, meta)
}
