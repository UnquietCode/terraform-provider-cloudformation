// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupInstancesDistribution() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"on_demand_allocation_strategy": {
				Type: schema.TypeString,
				Required: false,
			},
			"on_demand_base_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"on_demand_percentage_above_base_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"spot_allocation_strategy": {
				Type: schema.TypeString,
				Required: false,
			},
			"spot_instance_pools": {
				Type: schema.TypeInt,
				Required: false,
			},
			"spot_max_price": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}