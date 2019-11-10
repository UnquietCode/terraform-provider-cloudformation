// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPolicyStepAdjustment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_interval_lower_bound": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"metric_interval_upper_bound": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"scaling_adjustment": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
	}
}