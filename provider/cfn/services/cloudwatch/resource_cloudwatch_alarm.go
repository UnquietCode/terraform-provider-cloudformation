// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudWatchAlarm() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudWatchAlarmExists,
		Read: resourceCloudWatchAlarmRead,
		Create: resourceCloudWatchAlarmCreate,
		Update: resourceCloudWatchAlarmUpdate,
		Delete: resourceCloudWatchAlarmDelete,
		CustomizeDiff: resourceCloudWatchAlarmCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"actions_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"alarm_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"alarm_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"alarm_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"comparison_operator": {
				Type: schema.TypeString,
				Required: true,
			},
			"datapoints_to_alarm": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"dimensions": {
				Type: schema.TypeList,
				Elem: propertyAlarmDimension(),
				Optional: true,
			},
			"evaluate_low_sample_count_percentile": {
				Type: schema.TypeString,
				Optional: true,
			},
			"evaluation_periods": {
				Type: schema.TypeInt,
				Required: true,
			},
			"extended_statistic": {
				Type: schema.TypeString,
				Optional: true,
			},
			"insufficient_data_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"metric_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"metrics": {
				Type: schema.TypeSet,
				Elem: propertyAlarmMetricDataQuery(),
				Optional: true,
			},
			"namespace": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ok_actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"statistic": {
				Type: schema.TypeString,
				Optional: true,
			},
			"threshold": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"threshold_metric_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"treat_missing_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"unit": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudWatchAlarmExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudWatchAlarmRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudWatch::Alarm", ResourceCloudWatchAlarm(), data, meta)
}

func resourceCloudWatchAlarmCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudWatch::Alarm", ResourceCloudWatchAlarm(), data, meta)
}

func resourceCloudWatchAlarmUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudWatch::Alarm", ResourceCloudWatchAlarm(), data, meta)
}

func resourceCloudWatchAlarmDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudWatch::Alarm", data, meta)
}

func resourceCloudWatchAlarmCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
