// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html

package applicationautoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPolicyStepScalingPolicyConfiguration(extras...string) *schema.Resource {
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
			"adjustment_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cooldown": {
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
			"step_adjustments": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyStepAdjustment(),
				Optional: true,
			},
		},
	}
}
