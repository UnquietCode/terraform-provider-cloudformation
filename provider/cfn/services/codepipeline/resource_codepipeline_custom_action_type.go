// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodePipelineCustomActionType() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodePipelineCustomActionTypeExists,
		Read: resourceCodePipelineCustomActionTypeRead,
		Create: resourceCodePipelineCustomActionTypeCreate,
		Update: resourceCodePipelineCustomActionTypeUpdate,
		Delete: resourceCodePipelineCustomActionTypeDelete,
		CustomizeDiff: resourceCodePipelineCustomActionTypeCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration_properties": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeConfigurationProperties(),
				Optional: true,
			},
			"input_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				MaxItems: 1,
			},
			"output_artifact_details": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				MaxItems: 1,
			},
			"the_provider": {
				Type: schema.TypeString,
				Required: true,
			},
			"settings": {
				Type: schema.TypeList,
				Elem: propertyCustomActionTypeSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCodePipelineCustomActionTypeExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodePipelineCustomActionTypeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::CustomActionType", ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::CustomActionType", ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::CustomActionType", ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::CustomActionType", data, meta)
}

func resourceCodePipelineCustomActionTypeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::CodePipeline::CustomActionType", data, meta)
}

