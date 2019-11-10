// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupMixedInstancesPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instances_distribution": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupInstancesDistribution(),
				Required: false,
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