// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingInstruction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_dynamic_scaling": {
				Type: schema.TypeBool,
				Required: false,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"predictive_scaling_max_capacity_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Required: true,
			},
			"scaling_policy_update_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"min_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"target_tracking_configurations": {
				Type: schema.TypeList,
				Elem: propertyTargetTrackingConfiguration(),
				Required: true,
			},
			"predictive_scaling_max_capacity_buffer": {
				Type: schema.TypeInt,
				Required: false,
			},
			"customized_load_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyCustomizedLoadMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"predefined_load_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyPredefinedLoadMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"scheduled_action_buffer_time": {
				Type: schema.TypeInt,
				Required: false,
			},
			"max_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"predictive_scaling_mode": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}