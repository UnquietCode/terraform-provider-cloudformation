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

func propertySpotOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"spot_instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_interruption_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"max_price": {
				Type: schema.TypeString,
				Required: false,
			},
			"block_duration_minutes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"valid_until": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}