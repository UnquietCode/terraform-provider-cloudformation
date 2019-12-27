// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html

package codepipeline

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codePipelinePipelineType string = "AWS::CodePipeline::Pipeline"

func DatasourceCodePipelinePipeline() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodePipelinePipelineRead,
		
		Schema: map[string]*schema.Schema{
			"artifact_store": {
				Type: schema.TypeList,
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
			"tags": misc.PropertyTags(),
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

func datasourceCodePipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codePipelinePipelineType, DatasourceCodePipelinePipeline(), data, meta)
}
