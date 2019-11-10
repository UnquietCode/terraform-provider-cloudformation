// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPolicyStepScalingPolicyConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adjustment_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"cooldown": {
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
			"step_adjustments": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyStepAdjustment(),
				Required: false,
			},
		},
	}
}