// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceIoTAnalyticsPipeline() *schema.Resource {
	return &schema.Resource{
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

func resourceIoTAnalyticsPipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTAnalyticsPipelineType, ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTAnalyticsPipelineType, data, meta)
}

func resourceIoTAnalyticsPipelineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTAnalyticsPipelineType, data, meta)
}
