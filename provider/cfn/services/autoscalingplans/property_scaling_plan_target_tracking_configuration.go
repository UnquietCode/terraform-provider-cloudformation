// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html

package autoscalingplans

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPlanTargetTrackingConfiguration(extras...string) *schema.Resource {
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
			"scale_out_cooldown": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"target_value": {
				Type: schema.TypeFloat,
				Required: true,
			},
			"predefined_scaling_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanPredefinedScalingMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"disable_scale_in": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"scale_in_cooldown": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"estimated_instance_warmup": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"customized_scaling_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanCustomizedScalingMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
