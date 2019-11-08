// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTargetCapacitySpecificationRequest() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_target_capacity_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"total_target_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"on_demand_target_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"spot_target_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}