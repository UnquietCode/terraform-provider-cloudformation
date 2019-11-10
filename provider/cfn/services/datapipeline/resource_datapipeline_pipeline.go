// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceDataPipelinePipelineCreate,
		Read:   resourceDataPipelinePipelineRead,
		Update: resourceDataPipelinePipelineUpdate,
		Delete: resourceDataPipelinePipelineDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func resourceDataPipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DataPipeline::Pipeline", ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DataPipeline::Pipeline", data, meta)
}