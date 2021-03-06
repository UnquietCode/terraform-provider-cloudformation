// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html

package autoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupMixedInstancesPolicy(extras...string) *schema.Resource {
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
			"instances_distribution": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupInstancesDistribution(),
				Optional: true,
				MaxItems: 1,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupLaunchTemplate(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}
