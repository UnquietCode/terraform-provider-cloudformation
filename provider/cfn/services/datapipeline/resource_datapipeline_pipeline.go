// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html

package datapipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDataPipelinePipeline() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDataPipelinePipelineExists,
		Read: resourceDataPipelinePipelineRead,
		Create: resourceDataPipelinePipelineCreate,
		Update: resourceDataPipelinePipelineUpdate,
		Delete: resourceDataPipelinePipelineDelete,
		CustomizeDiff: resourceDataPipelinePipelineCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameter_objects": {
				Type: schema.TypeList,
				Elem: propertyPipelineParameterObject(),
				Required: true,
			},
			"parameter_values": {
				Type: schema.TypeList,
				Elem: propertyPipelineParameterValue(),
				Optional: true,
			},
			"pipeline_objects": {
				Type: schema.TypeList,
				Elem: propertyPipelinePipelineObject(),
				Optional: true,
			},
			"pipeline_tags": {
				Type: schema.TypeList,
				Elem: propertyPipelinePipelineTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDataPipelinePipelineExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDataPipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DataPipeline::Pipeline", data, meta)
}

func resourceDataPipelinePipelineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::DataPipeline::Pipeline", data, meta)
}

