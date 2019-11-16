// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingScheduledActionType string = "AWS::AutoScaling::ScheduledAction"

var autoScalingScheduledActionProperties map[string]string = map[string]string{
	"auto_scaling_group_name": "AutoScalingGroupName",
	"desired_capacity": "DesiredCapacity",
	"end_time": "EndTime",
	"max_size": "MaxSize",
	"min_size": "MinSize",
	"recurrence": "Recurrence",
	"start_time": "StartTime",
}

func ResourceAutoScalingScheduledAction() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAutoScalingScheduledActionExists,
		Read: resourceAutoScalingScheduledActionRead,
		Create: resourceAutoScalingScheduledActionCreate,
		Update: resourceAutoScalingScheduledActionUpdate,
		Delete: resourceAutoScalingScheduledActionDelete,
		CustomizeDiff: resourceAutoScalingScheduledActionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"desired_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"end_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"max_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"min_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"recurrence": {
				Type: schema.TypeString,
				Optional: true,
			},
			"start_time": {
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

func resourceAutoScalingScheduledActionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAutoScalingScheduledActionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, meta)
}

func resourceAutoScalingScheduledActionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, autoScalingScheduledActionProperties, meta)
}

func resourceAutoScalingScheduledActionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, autoScalingScheduledActionProperties, meta)
}

func resourceAutoScalingScheduledActionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(autoScalingScheduledActionType, data, meta)
}

func resourceAutoScalingScheduledActionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(autoScalingScheduledActionType, data, meta)
}
