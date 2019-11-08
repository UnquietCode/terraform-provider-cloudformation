// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package budgets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCostTypes() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"include_support": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_other_subscription": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_tax": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_subscription": {
				Type: schema.TypeBool,
				Required: false,
			},
			"use_blended": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_upfront": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_discount": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_credit": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_recurring": {
				Type: schema.TypeBool,
				Required: false,
			},
			"use_amortized": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_refund": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}