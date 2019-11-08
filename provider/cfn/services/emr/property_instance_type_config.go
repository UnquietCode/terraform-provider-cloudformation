// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceTypeConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bid_price": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"bid_price_as_percentage_of_on_demand_price": {
				Type: schema.TypeFloat,
				Required: false,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyConfiguration(),
				Required: false,
				ForceNew: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyEbsConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"weighted_capacity": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
		},
	}
}