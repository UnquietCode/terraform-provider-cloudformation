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

func ResourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Schema: map[string]*schema.Schema{
			"artifact_store": {
				Type: schema.TypeList,
				Elem: propertyPipelineArtifactStore(),
				Required: false,
				MaxItems: 1,
			},
			"artifact_stores": {
				Type: schema.TypeSet,
				Elem: propertyPipelineArtifactStoreMap(),
				Required: false,
			},
			"disable_inbound_stage_transitions": {
				Type: schema.TypeSet,
				Elem: propertyPipelineStageTransition(),
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
				Elem: propertyPipelineStageDeclaration(),
				Required: true,
			},
		},
	}
}

func resourcePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::Pipeline", data, meta)
}

func resourcePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::Pipeline", data, meta)
}

func resourcePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::Pipeline", data, meta)
}

func resourcePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::Pipeline", data, meta)
}