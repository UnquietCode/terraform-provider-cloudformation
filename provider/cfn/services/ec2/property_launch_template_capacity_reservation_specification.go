// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchTemplateCapacityReservationSpecification(extras...string) *schema.Resource {
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
			"capacity_reservation_preference": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateCapacityReservationPreference(),
				Required: false,
				MaxItems: 1,
			},
			"capacity_reservation_target": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateCapacityReservationTarget(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}