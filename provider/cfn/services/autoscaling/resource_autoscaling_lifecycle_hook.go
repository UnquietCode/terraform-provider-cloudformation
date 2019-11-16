// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingLifecycleHookType string = "AWS::AutoScaling::LifecycleHook"

var autoScalingLifecycleHookProperties map[string]string = map[string]string{
	"auto_scaling_group_name": "AutoScalingGroupName",
	"default_result": "DefaultResult",
	"heartbeat_timeout": "HeartbeatTimeout",
	"lifecycle_hook_name": "LifecycleHookName",
	"lifecycle_transition": "LifecycleTransition",
	"notification_metadata": "NotificationMetadata",
	"notification_target_arn": "NotificationTargetARN",
	"role_arn": "RoleARN",
}

func ResourceAutoScalingLifecycleHook() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAutoScalingLifecycleHookExists,
		Read: resourceAutoScalingLifecycleHookRead,
		Create: resourceAutoScalingLifecycleHookCreate,
		Update: resourceAutoScalingLifecycleHookUpdate,
		Delete: resourceAutoScalingLifecycleHookDelete,
		CustomizeDiff: resourceAutoScalingLifecycleHookCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_result": {
				Type: schema.TypeString,
				Optional: true,
			},
			"heartbeat_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"lifecycle_hook_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"lifecycle_transition": {
				Type: schema.TypeString,
				Required: true,
			},
			"notification_metadata": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_target_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
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

func resourceAutoScalingLifecycleHookExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAutoScalingLifecycleHookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingLifecycleHookType, ResourceAutoScalingLifecycleHook(), data, meta)
}

func resourceAutoScalingLifecycleHookCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(autoScalingLifecycleHookType, ResourceAutoScalingLifecycleHook(), data, autoScalingLifecycleHookProperties, meta)
}

func resourceAutoScalingLifecycleHookUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(autoScalingLifecycleHookType, ResourceAutoScalingLifecycleHook(), data, autoScalingLifecycleHookProperties, meta)
}

func resourceAutoScalingLifecycleHookDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(autoScalingLifecycleHookType, data, meta)
}

func resourceAutoScalingLifecycleHookCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(autoScalingLifecycleHookType, data, meta)
}
