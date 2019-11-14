// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogsMetricFilter() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLogsMetricFilterExists,
		Read: resourceLogsMetricFilterRead,
		Create: resourceLogsMetricFilterCreate,
		Update: resourceLogsMetricFilterUpdate,
		Delete: resourceLogsMetricFilterDelete,
		CustomizeDiff: resourceLogsMetricFilterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"filter_pattern": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_transformations": {
				Type: schema.TypeSet,
				Elem: propertyMetricFilterMetricTransformation(),
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

func resourceLogsMetricFilterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLogsMetricFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::MetricFilter", ResourceLogsMetricFilter(), data, meta)
}

func resourceLogsMetricFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::MetricFilter", ResourceLogsMetricFilter(), data, meta)
}

func resourceLogsMetricFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::MetricFilter", ResourceLogsMetricFilter(), data, meta)
}

func resourceLogsMetricFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::MetricFilter", data, meta)
}

func resourceLogsMetricFilterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

