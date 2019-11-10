// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceLogsMetricFilterCreate,
		Read:   resourceLogsMetricFilterRead,
		Update: resourceLogsMetricFilterUpdate,
		Delete: resourceLogsMetricFilterDelete,

		Schema: map[string]*schema.Schema{
			"filter_pattern": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metric_transformations": {
				Type: schema.TypeSet,
				Elem: propertyMetricFilterMetricTransformation(),
				Required: true,
			},
		},
	}
}

func resourceLogsMetricFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::MetricFilter", data, meta)
}

func resourceLogsMetricFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::MetricFilter", data, meta)
}

func resourceLogsMetricFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::MetricFilter", data, meta)
}

func resourceLogsMetricFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::MetricFilter", data, meta)
}