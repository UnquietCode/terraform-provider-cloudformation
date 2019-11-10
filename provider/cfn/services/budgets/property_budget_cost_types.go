// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package budgets

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBudgetCostTypes(extras...string) *schema.Resource {
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