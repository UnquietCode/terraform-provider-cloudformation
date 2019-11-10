// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
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
				Required: false,
			},
			"pipeline_objects": {
				Type: schema.TypeList,
				Elem: propertyPipelinePipelineObject(),
				Required: false,
			},
			"pipeline_tags": {
				Type: schema.TypeList,
				Elem: propertyPipelinePipelineTag(),
				Required: false,
			},
		},
	}
}

func resourceDataPipelinePipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DataPipeline::Pipeline", data, meta)
}

func resourceDataPipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DataPipeline::Pipeline", data, meta)
}

func resourceDataPipelinePipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DataPipeline::Pipeline", data, meta)
}

func resourceDataPipelinePipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DataPipeline::Pipeline", data, meta)
}