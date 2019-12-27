// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html

package datapipeline

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dataPipelinePipelineType string = "AWS::DataPipeline::Pipeline"

func DatasourceDataPipelinePipeline() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDataPipelinePipelineRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceDataPipelinePipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dataPipelinePipelineType, DatasourceDataPipelinePipeline(), data, meta)
}
