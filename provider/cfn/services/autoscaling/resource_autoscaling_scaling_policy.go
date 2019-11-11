// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingScalingPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAutoScalingScalingPolicyExists,
		Read: resourceAutoScalingScalingPolicyRead,
		Create: resourceAutoScalingScalingPolicyCreate,
		Update: resourceAutoScalingScalingPolicyUpdate,
		Delete: resourceAutoScalingScalingPolicyDelete,
		CustomizeDiff: resourceAutoScalingScalingPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"adjustment_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cooldown": {
				Type: schema.TypeString,
				Optional: true,
			},
			"estimated_instance_warmup": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"metric_aggregation_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"min_adjustment_magnitude": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"policy_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scaling_adjustment": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"step_adjustments": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyStepAdjustment(),
				Optional: true,
			},
			"target_tracking_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyTargetTrackingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingScalingPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAutoScalingScalingPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::ScalingPolicy", ResourceAutoScalingScalingPolicy(), data, meta)
}

func resourceAutoScalingScalingPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::ScalingPolicy", ResourceAutoScalingScalingPolicy(), data, meta)
}

func resourceAutoScalingScalingPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::ScalingPolicy", ResourceAutoScalingScalingPolicy(), data, meta)
}

func resourceAutoScalingScalingPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::ScalingPolicy", data, meta)
}

func resourceAutoScalingScalingPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::AutoScaling::ScalingPolicy", data, meta)
}

