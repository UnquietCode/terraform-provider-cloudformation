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

func ResourceScalingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceScalingPolicyCreate,
		Read:   resourceScalingPolicyRead,
		Update: resourceScalingPolicyUpdate,
		Delete: resourceScalingPolicyDelete,

		Schema: map[string]*schema.Schema{
			"adjustment_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cooldown": {
				Type: schema.TypeString,
				Required: false,
			},
			"estimated_instance_warmup": {
				Type: schema.TypeInt,
				Required: false,
			},
			"metric_aggregation_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"min_adjustment_magnitude": {
				Type: schema.TypeInt,
				Required: false,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"scaling_adjustment": {
				Type: schema.TypeInt,
				Required: false,
			},
			"step_adjustments": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyStepAdjustment(),
				Required: false,
			},
			"target_tracking_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyTargetTrackingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceScalingPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::ScalingPolicy", data, meta)
}