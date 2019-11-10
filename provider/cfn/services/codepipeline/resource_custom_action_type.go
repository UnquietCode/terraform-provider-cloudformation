// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCustomActionType() *schema.Resource {
	return &schema.Resource{
		Create: resourceCustomActionTypeCreate,
		Read:   resourceCustomActionTypeRead,
		Update: resourceCustomActionTypeUpdate,
		Delete: resourceCustomActionTypeDelete,

		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration_properties": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeConfigurationProperties(),
				Required: false,
				ForceNew: true,
			},
			"input_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"output_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"provider": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"settings": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeSettings(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCustomActionTypeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCustomActionTypeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCustomActionTypeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCustomActionTypeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::CustomActionType", data, meta)
}