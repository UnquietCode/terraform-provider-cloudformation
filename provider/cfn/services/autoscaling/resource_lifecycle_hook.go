// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLifecycleHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceLifecycleHookCreate,
		Read:   resourceLifecycleHookRead,
		Update: resourceLifecycleHookUpdate,
		Delete: resourceLifecycleHookDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_result": {
				Type: schema.TypeString,
				Required: false,
			},
			"heartbeat_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"lifecycle_hook_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"lifecycle_transition": {
				Type: schema.TypeString,
				Required: true,
			},
			"notification_metadata": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_target_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceLifecycleHookCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::LifecycleHook", data, meta)
}

func resourceLifecycleHookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::LifecycleHook", data, meta)
}

func resourceLifecycleHookUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::LifecycleHook", data, meta)
}

func resourceLifecycleHookDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::LifecycleHook", data, meta)
}