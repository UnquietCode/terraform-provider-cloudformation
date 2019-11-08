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

func ResourceCodePipelinePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodePipelinePipelineCreate,
		Read:   resourceCodePipelinePipelineRead,
		Update: resourceCodePipelinePipelineUpdate,
		Delete: resourceCodePipelinePipelineDelete,

		Schema: map[string]*schema.Schema{
			"artifact_store": {
				Type: schema.TypeList,
				Elem: propertyArtifactStore(),
				Required: false,
				MaxItems: 1,
			},
			"artifact_stores": {
				Type: schema.TypeSet,
				Elem: propertyArtifactStoreMap(),
				Required: false,
			},
			"disable_inbound_stage_transitions": {
				Type: schema.TypeSet,
				Elem: propertyStageTransition(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"restart_execution_on_update": {
				Type: schema.TypeBool,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"stages": {
				Type: schema.TypeSet,
				Elem: propertyStageDeclaration(),
				Required: true,
			},
		},
	}
}

func resourceCodePipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::Pipeline", data, meta)
}

func resourceCodePipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::Pipeline", data, meta)
}

func resourceCodePipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::Pipeline", data, meta)
}

func resourceCodePipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::Pipeline", data, meta)
}