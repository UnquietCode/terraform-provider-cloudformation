// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html

package autoscalingplans

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPlanScalingInstruction(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_dynamic_scaling": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"predictive_scaling_max_capacity_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Required: true,
			},
			"scaling_policy_update_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"min_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"target_tracking_configurations": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanTargetTrackingConfiguration(),
				Required: true,
			},
			"predictive_scaling_max_capacity_buffer": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"customized_load_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanCustomizedLoadMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"predefined_load_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanPredefinedLoadMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"scheduled_action_buffer_time": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"max_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"predictive_scaling_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

