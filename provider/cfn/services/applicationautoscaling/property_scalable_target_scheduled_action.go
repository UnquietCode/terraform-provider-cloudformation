// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html

package applicationautoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalableTargetScheduledAction(extras...string) *schema.Resource {
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
			"end_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scalable_target_action": {
				Type: schema.TypeList,
				Elem: propertyScalableTargetScalableTargetAction(),
				Optional: true,
				MaxItems: 1,
			},
			"schedule": {
				Type: schema.TypeString,
				Required: true,
			},
			"scheduled_action_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"start_time": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
