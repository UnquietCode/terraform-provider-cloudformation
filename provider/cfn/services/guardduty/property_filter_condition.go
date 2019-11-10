// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFilterCondition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lt": {
				Type: schema.TypeInt,
				Required: false,
			},
			"gte": {
				Type: schema.TypeInt,
				Required: false,
			},
			"neq": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"eq": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"lte": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}