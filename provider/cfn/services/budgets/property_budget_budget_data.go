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

func propertyBudgetBudgetData(extras...string) *schema.Resource {
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
			"budget_limit": {
				Type: schema.TypeList,
				Elem: propertyBudgetSpend(),
				Required: false,
				MaxItems: 1,
			},
			"time_period": {
				Type: schema.TypeList,
				Elem: propertyBudgetTimePeriod(),
				Required: false,
				MaxItems: 1,
			},
			"time_unit": {
				Type: schema.TypeString,
				Required: true,
			},
			"planned_budget_limits": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"cost_filters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"budget_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cost_types": {
				Type: schema.TypeList,
				Elem: propertyBudgetCostTypes(),
				Required: false,
				MaxItems: 1,
			},
			"budget_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}