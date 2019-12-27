// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-pipeline.html

package iotanalytics

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTAnalyticsPipelineType string = "AWS::IoTAnalytics::Pipeline"

func DatasourceIoTAnalyticsPipeline() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIoTAnalyticsPipelineRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceIoTAnalyticsPipelineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsPipelineType, DatasourceIoTAnalyticsPipeline(), data, meta)
}
