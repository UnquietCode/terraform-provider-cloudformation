// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterInstanceFleetConfig(extras...string) *schema.Resource {
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
			"instance_type_configs": {
				Type: schema.TypeSet,
				Elem: propertyClusterInstanceTypeConfig(),
				Required: false,
				ForceNew: true,
			},
			"launch_specifications": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceFleetProvisioningSpecifications(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"target_on_demand_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"target_spot_capacity": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}