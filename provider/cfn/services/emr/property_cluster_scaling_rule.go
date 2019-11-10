// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterScalingRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList,
				Elem: propertyClusterScalingAction(),
				Required: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"trigger": {
				Type: schema.TypeList,
				Elem: propertyClusterScalingTrigger(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}