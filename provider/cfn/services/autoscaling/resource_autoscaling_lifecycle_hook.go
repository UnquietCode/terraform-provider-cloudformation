// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingLifecycleHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutoScalingLifecycleHookCreate,
		Read:   resourceAutoScalingLifecycleHookRead,
		Update: resourceAutoScalingLifecycleHookUpdate,
		Delete: resourceAutoScalingLifecycleHookDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingLifecycleHookCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::LifecycleHook", ResourceAutoScalingLifecycleHook(), data, meta)
}

func resourceAutoScalingLifecycleHookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::LifecycleHook", ResourceAutoScalingLifecycleHook(), data, meta)
}

func resourceAutoScalingLifecycleHookUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::LifecycleHook", ResourceAutoScalingLifecycleHook(), data, meta)
}

func resourceAutoScalingLifecycleHookDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::LifecycleHook", data, meta)
}