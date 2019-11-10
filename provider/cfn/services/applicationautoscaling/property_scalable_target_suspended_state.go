// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html

package applicationautoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalableTargetSuspendedState(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
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