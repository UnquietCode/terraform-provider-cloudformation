// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAlarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmCreate,
		Read:   resourceAlarmRead,
		Update: resourceAlarmUpdate,
		Delete: resourceAlarmDelete,

		Schema: map[string]*schema.Schema{
			"actions_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"alarm_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"alarm_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"alarm_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"comparison_operator": {
				Type: schema.TypeString,
				Required: true,
			},
			"datapoints_to_alarm": {
				Type: schema.TypeInt,
				Required: false,
			},
			"dimensions": {
				Type: schema.TypeList,
				Elem: propertyAlarmDimension(),
				Required: false,
			},
			"evaluate_low_sample_count_percentile": {
				Type: schema.TypeString,
				Required: false,
			},
			"evaluation_periods": {
				Type: schema.TypeInt,
				Required: true,
			},
			"extended_statistic": {
				Type: schema.TypeString,
				Required: false,
			},
			"insufficient_data_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"metrics": {
				Type: schema.TypeSet,
				Elem: propertyAlarmMetricDataQuery(),
				Required: false,
			},
			"namespace": {
				Type: schema.TypeString,
				Required: false,
			},
			"ok_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"statistic": {
				Type: schema.TypeString,
				Required: false,
			},
			"threshold": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"threshold_metric_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"treat_missing_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"unit": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceAlarmCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudWatch::Alarm", data, meta)
}

func resourceAlarmRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudWatch::Alarm", data, meta)
}

func resourceAlarmUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudWatch::Alarm", data, meta)
}

func resourceAlarmDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudWatch::Alarm", data, meta)
}