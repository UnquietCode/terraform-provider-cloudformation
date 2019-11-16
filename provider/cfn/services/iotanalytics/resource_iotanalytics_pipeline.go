// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-pipeline.html

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTAnalyticsPipelineType string = "AWS::IoTAnalytics::Pipeline"

var ioTAnalyticsPipelineProperties map[string]string = map[string]string{
	"pipeline_name": "PipelineName",
	"tags": "Tags",
	"pipeline_activities": "PipelineActivities",
}

func ResourceIoTAnalyticsPipeline() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTAnalyticsPipelineExists,
		Read: resourceIoTAnalyticsPipelineRead,
		Create: resourceIoTAnalyticsPipelineCreate,
		Update: resourceIoTAnalyticsPipelineUpdate,
		Delete: resourceIoTAnalyticsPipelineDelete,
		CustomizeDiff: resourceIoTAnalyticsPipelineCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"pipeline_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"pipeline_activities": {
				Type: schema.TypeList,
				Elem: propertyPipelineActivity(),
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

func resourceIoTAnalyticsPipelineExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTAnalyticsPipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, ioTAnalyticsPipelineProperties, meta)
}

func resourceIoTAnalyticsPipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, ioTAnalyticsPipelineProperties, meta)
}

func resourceIoTAnalyticsPipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTAnalyticsPipelineType, data, meta)
}

func resourceIoTAnalyticsPipelineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTAnalyticsPipelineType, data, meta)
}
