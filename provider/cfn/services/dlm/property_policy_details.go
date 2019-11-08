// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dlm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

func propertyPolicyDetails() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"schedules": {
				Type: schema.TypeList,
				Elem: propertySchedule(),
				Required: false,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeList,
				Elem: propertyParameters(),
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