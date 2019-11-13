// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceIoTAnalyticsPipeline() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTAnalyticsPipelineExists,
		Read:   resourceIoTAnalyticsPipelineRead,
		Create: resourceIoTAnalyticsPipelineCreate,
		Update: resourceIoTAnalyticsPipelineUpdate,
		Delete: resourceIoTAnalyticsPipelineDelete,
		
		Schema: map[string]*schema.Schema{
			"pipeline_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
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
	return plugin.ResourceRead("AWS::IoTAnalytics::Pipeline", ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Pipeline", ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Pipeline", ResourceIoTAnalyticsPipeline(), data, meta)
}

func resourceIoTAnalyticsPipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Pipeline", data, meta)
}