// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dlm

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

func propertyLifecyclePolicyPolicyDetails(extras...string) *schema.Resource {
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
			"resource_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"schedules": {
				Type: schema.TypeList,
				Elem: propertyLifecyclePolicySchedule(),
				Required: false,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeList,
				Elem: propertyLifecyclePolicyParameters(),
				Required: false,
				MaxItems: 1,
			},
			"target_tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}