// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeBuildProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeBuildProjectCreate,
		Read:   resourceCodeBuildProjectRead,
		Update: resourceCodeBuildProjectUpdate,
		Delete: resourceCodeBuildProjectDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyVpcConfig(),
				Required: false,
				MaxItems: 1,
			},
			"secondary_sources": {
				Type: schema.TypeList,
				Elem: propertySource(),
				Required: false,
			},
			"encryption_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"triggers": {
				Type: schema.TypeList,
				Elem: propertyProjectTriggers(),
				Required: false,
				MaxItems: 1,
			},
			"secondary_artifacts": {
				Type: schema.TypeList,
				Elem: propertyArtifacts(),
				Required: false,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertySource(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"artifacts": {
				Type: schema.TypeList,
				Elem: propertyArtifacts(),
				Required: true,
				MaxItems: 1,
			},
			"badge_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"logs_config": {
				Type: schema.TypeList,
				Elem: propertyLogsConfig(),
				Required: false,
				MaxItems: 1,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"queued_timeout_in_minutes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyEnvironment(),
				Required: true,
				MaxItems: 1,
			},
			"secondary_source_versions": {
				Type: schema.TypeList,
				Elem: propertyProjectSourceVersion(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"timeout_in_minutes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"cache": {
				Type: schema.TypeList,
				Elem: propertyProjectCache(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceCodeBuildProjectCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeBuild::Project", data, meta)
}

func resourceCodeBuildProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeBuild::Project", data, meta)
}

func resourceCodeBuildProjectUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeBuild::Project", data, meta)
}

func resourceCodeBuildProjectDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeBuild::Project", data, meta)
}