// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodePipelineCustomActionType() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodePipelineCustomActionTypeCreate,
		Read:   resourceCodePipelineCustomActionTypeRead,
		Update: resourceCodePipelineCustomActionTypeUpdate,
		Delete: resourceCodePipelineCustomActionTypeDelete,

		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration_properties": {
				Type: schema.TypeSet,
				Elem: propertyConfigurationProperties(),
				Required: false,
				ForceNew: true,
			},
			"input_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyArtifactDetails(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"output_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyArtifactDetails(),
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
				Elem: propertySettings(),
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

func resourceCodePipelineCustomActionTypeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCodePipelineCustomActionTypeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCodePipelineCustomActionTypeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCodePipelineCustomActionTypeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::CustomActionType", data, meta)
}