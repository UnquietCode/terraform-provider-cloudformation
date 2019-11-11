// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplicationAutoScalingScalingPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApplicationAutoScalingScalingPolicyExists,
		Read: resourceApplicationAutoScalingScalingPolicyRead,
		Create: resourceApplicationAutoScalingScalingPolicyCreate,
		Update: resourceApplicationAutoScalingScalingPolicyUpdate,
		Delete: resourceApplicationAutoScalingScalingPolicyDelete,
		CustomizeDiff: resourceApplicationAutoScalingScalingPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scaling_target_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Optional: true,
			},
			"step_scaling_policy_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyStepScalingPolicyConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"target_tracking_scaling_policy_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyTargetTrackingScalingPolicyConfiguration(),
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

func resourceApplicationAutoScalingScalingPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApplicationAutoScalingScalingPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApplicationAutoScaling::ScalingPolicy", ResourceApplicationAutoScalingScalingPolicy(), data, meta)
}

func resourceApplicationAutoScalingScalingPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApplicationAutoScaling::ScalingPolicy", ResourceApplicationAutoScalingScalingPolicy(), data, meta)
}

func resourceApplicationAutoScalingScalingPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApplicationAutoScaling::ScalingPolicy", ResourceApplicationAutoScalingScalingPolicy(), data, meta)
}

func resourceApplicationAutoScalingScalingPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}

func resourceApplicationAutoScalingScalingPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}

