// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMetricFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetricFilterCreate,
		Read:   resourceMetricFilterRead,
		Update: resourceMetricFilterUpdate,
		Delete: resourceMetricFilterDelete,

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

func resourceMetricFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::MetricFilter", data, meta)
}

func resourceMetricFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::MetricFilter", data, meta)
}

func resourceMetricFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::MetricFilter", data, meta)
}

func resourceMetricFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::MetricFilter", data, meta)
}