// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codePipelinePipelineType string = "AWS::CodePipeline::Pipeline"

var codePipelinePipelineProperties map[string]string = map[string]string{
	"artifact_store": "ArtifactStore",
	"artifact_stores": "ArtifactStores",
	"disable_inbound_stage_transitions": "DisableInboundStageTransitions",
	"name": "Name",
	"restart_execution_on_update": "RestartExecutionOnUpdate",
	"role_arn": "RoleArn",
	"stages": "Stages",
}

func ResourceCodePipelinePipeline() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodePipelinePipelineExists,
		Read: resourceCodePipelinePipelineRead,
		Create: resourceCodePipelinePipelineCreate,
		Update: resourceCodePipelinePipelineUpdate,
		Delete: resourceCodePipelinePipelineDelete,
		CustomizeDiff: resourceCodePipelinePipelineCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"artifact_store": {
				Type: schema.TypeSet,
				Elem: propertyPipelineArtifactStore(),
				Optional: true,
				MaxItems: 1,
			},
			"artifact_stores": {
				Type: schema.TypeSet,
				Elem: propertyPipelineArtifactStoreMap(),
				Optional: true,
			},
			"disable_inbound_stage_transitions": {
				Type: schema.TypeSet,
				Elem: propertyPipelineStageTransition(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"restart_execution_on_update": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"stages": {
				Type: schema.TypeSet,
				Elem: propertyPipelineStageDeclaration(),
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

func resourceCodePipelinePipelineExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodePipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codePipelinePipelineType, ResourceCodePipelinePipeline(), data, meta)
}

func resourceCodePipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codePipelinePipelineType, ResourceCodePipelinePipeline(), data, codePipelinePipelineProperties, meta)
}

func resourceCodePipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codePipelinePipelineType, ResourceCodePipelinePipeline(), data, codePipelinePipelineProperties, meta)
}

func resourceCodePipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codePipelinePipelineType, data, meta)
}

func resourceCodePipelinePipelineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codePipelinePipelineType, data, meta)
}
