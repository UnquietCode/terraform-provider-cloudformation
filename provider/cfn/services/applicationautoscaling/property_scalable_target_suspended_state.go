// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalableTargetSuspendedState() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dynamic_scaling_in_suspended": {
				Type: schema.TypeBool,
				Required: false,
			},
			"dynamic_scaling_out_suspended": {
				Type: schema.TypeBool,
				Required: false,
			},
			"scheduled_scaling_suspended": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}