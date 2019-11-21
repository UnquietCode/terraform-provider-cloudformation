// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html

package autoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingScheduledActionType string = "AWS::AutoScaling::ScheduledAction"

func ResourceAutoScalingScheduledAction() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAutoScalingScheduledActionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, meta)
}

func resourceAutoScalingScheduledActionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, meta)
}

func resourceAutoScalingScheduledActionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(autoScalingScheduledActionType, ResourceAutoScalingScheduledAction(), data, meta)
}

func resourceAutoScalingScheduledActionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(autoScalingScheduledActionType, data, meta)
}

func resourceAutoScalingScheduledActionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(autoScalingScheduledActionType, data, meta)
}
