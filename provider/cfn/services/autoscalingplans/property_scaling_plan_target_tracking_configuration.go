// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPlanTargetTrackingConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scale_out_cooldown": {
				Type: schema.TypeInt,
				Required: false,
			},
			"target_value": {
				Type: schema.TypeFloat,
				Required: true,
			},
			"predefined_scaling_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanPredefinedScalingMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"disable_scale_in": {
				Type: schema.TypeBool,
				Required: false,
			},
			"scale_in_cooldown": {
				Type: schema.TypeInt,
				Required: false,
			},
			"estimated_instance_warmup": {
				Type: schema.TypeInt,
				Required: false,
			},
			"customized_scaling_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanCustomizedScalingMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}