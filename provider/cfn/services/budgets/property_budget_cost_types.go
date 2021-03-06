// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"include_support": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_other_subscription": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_tax": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_subscription": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"use_blended": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_upfront": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_discount": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_credit": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_recurring": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"use_amortized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_refund": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
