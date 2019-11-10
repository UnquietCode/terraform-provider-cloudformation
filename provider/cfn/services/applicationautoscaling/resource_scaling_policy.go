// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package applicationautoscaling

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
			"policy_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scaling_target_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"step_scaling_policy_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyStepScalingPolicyConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"target_tracking_scaling_policy_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyTargetTrackingScalingPolicyConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceScalingPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}

func resourceScalingPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApplicationAutoScaling::ScalingPolicy", data, meta)
}