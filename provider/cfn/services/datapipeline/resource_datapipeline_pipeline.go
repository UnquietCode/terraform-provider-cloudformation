// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html

package datapipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dataPipelinePipelineType string = "AWS::DataPipeline::Pipeline"

func ResourceDataPipelinePipeline() *schema.Resource {
	return &schema.Resource{
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

func resourceDataPipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dataPipelinePipelineType, ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dataPipelinePipelineType, ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dataPipelinePipelineType, ResourceDataPipelinePipeline(), data, meta)
}

func resourceDataPipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dataPipelinePipelineType, data, meta)
}

func resourceDataPipelinePipelineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dataPipelinePipelineType, data, meta)
}
