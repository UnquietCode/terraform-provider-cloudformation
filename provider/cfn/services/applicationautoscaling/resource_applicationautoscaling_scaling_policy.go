// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const applicationAutoScalingScalingPolicyType string = "AWS::ApplicationAutoScaling::ScalingPolicy"

var applicationAutoScalingScalingPolicyProperties map[string]string = map[string]string{
	"policy_name": "PolicyName",
	"policy_type": "PolicyType",
	"resource_id": "ResourceId",
	"scalable_dimension": "ScalableDimension",
	"scaling_target_id": "ScalingTargetId",
	"service_namespace": "ServiceNamespace",
	"step_scaling_policy_configuration": "StepScalingPolicyConfiguration",
	"target_tracking_scaling_policy_configuration": "TargetTrackingScalingPolicyConfiguration",
}

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
	return plugin.ResourceRead(applicationAutoScalingScalingPolicyType, ResourceApplicationAutoScalingScalingPolicy(), data, meta)
}

func resourceApplicationAutoScalingScalingPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(applicationAutoScalingScalingPolicyType, ResourceApplicationAutoScalingScalingPolicy(), data, applicationAutoScalingScalingPolicyProperties, meta)
}

func resourceApplicationAutoScalingScalingPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(applicationAutoScalingScalingPolicyType, ResourceApplicationAutoScalingScalingPolicy(), data, applicationAutoScalingScalingPolicyProperties, meta)
}

func resourceApplicationAutoScalingScalingPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(applicationAutoScalingScalingPolicyType, data, meta)
}

func resourceApplicationAutoScalingScalingPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(applicationAutoScalingScalingPolicyType, data, meta)
}
