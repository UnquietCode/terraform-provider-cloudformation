// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTAnalyticsPipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTAnalyticsPipelineCreate,
		Read:   resourceIoTAnalyticsPipelineRead,
		Update: resourceIoTAnalyticsPipelineUpdate,
		Delete: resourceIoTAnalyticsPipelineDelete,

		Schema: map[string]*schema.Schema{
			"pipeline_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"pipeline_activities": {
				Type: schema.TypeList,
				Elem: propertyActivity(),
				Required: true,
			},
		},
	}
}

func resourceIoTAnalyticsPipelineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Pipeline", data, meta)
}

func resourceIoTAnalyticsPipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTAnalytics::Pipeline", data, meta)
}

func resourceIoTAnalyticsPipelineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Pipeline", data, meta)
}

func resourceIoTAnalyticsPipelineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Pipeline", data, meta)
}