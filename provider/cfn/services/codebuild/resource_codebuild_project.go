// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeBuildProjectType string = "AWS::CodeBuild::Project"

var codeBuildProjectProperties map[string]string = map[string]string{
	"description": "Description",
	"vpc_config": "VpcConfig",
	"secondary_sources": "SecondarySources",
	"encryption_key": "EncryptionKey",
	"source_version": "SourceVersion",
	"triggers": "Triggers",
	"secondary_artifacts": "SecondaryArtifacts",
	"source": "Source",
	"name": "Name",
	"artifacts": "Artifacts",
	"badge_enabled": "BadgeEnabled",
	"logs_config": "LogsConfig",
	"service_role": "ServiceRole",
	"queued_timeout_in_minutes": "QueuedTimeoutInMinutes",
	"environment": "Environment",
	"secondary_source_versions": "SecondarySourceVersions",
	"tags": "Tags",
	"timeout_in_minutes": "TimeoutInMinutes",
	"cache": "Cache",
}

func ResourceCodeBuildProject() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodeBuildProjectExists,
		Read: resourceCodeBuildProjectRead,
		Create: resourceCodeBuildProjectCreate,
		Update: resourceCodeBuildProjectUpdate,
		Delete: resourceCodeBuildProjectDelete,
		CustomizeDiff: resourceCodeBuildProjectCustomizeDiff,
		
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
			},
		},
	}
}

func resourceCodeBuildProjectExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodeBuildProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeBuildProjectType, ResourceCodeBuildProject(), data, meta)
}

func resourceCodeBuildProjectCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeBuildProjectType, ResourceCodeBuildProject(), data, codeBuildProjectProperties, meta)
}

func resourceCodeBuildProjectUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeBuildProjectType, ResourceCodeBuildProject(), data, codeBuildProjectProperties, meta)
}

func resourceCodeBuildProjectDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeBuildProjectType, data, meta)
}

func resourceCodeBuildProjectCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeBuildProjectType, data, meta)
}
