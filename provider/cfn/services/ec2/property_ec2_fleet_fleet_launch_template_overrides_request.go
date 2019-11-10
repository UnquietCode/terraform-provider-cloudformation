// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEC2FleetFleetLaunchTemplateOverridesRequest() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"weighted_capacity": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"priority": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"max_price": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}