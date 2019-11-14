// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html

package autoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupInstancesDistribution(extras...string) *schema.Resource {
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
			"on_demand_allocation_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"on_demand_base_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"on_demand_percentage_above_base_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"spot_allocation_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"spot_instance_pools": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"spot_max_price": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

