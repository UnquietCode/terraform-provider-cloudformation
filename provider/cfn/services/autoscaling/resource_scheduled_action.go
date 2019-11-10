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

func ResourceScheduledAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceScheduledActionCreate,
		Read:   resourceScheduledActionRead,
		Update: resourceScheduledActionUpdate,
		Delete: resourceScheduledActionDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"desired_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"end_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"max_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"min_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"recurrence": {
				Type: schema.TypeString,
				Required: false,
			},
			"start_time": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceScheduledActionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::ScheduledAction", data, meta)
}

func resourceScheduledActionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::ScheduledAction", data, meta)
}

func resourceScheduledActionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::ScheduledAction", data, meta)
}

func resourceScheduledActionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::ScheduledAction", data, meta)
}