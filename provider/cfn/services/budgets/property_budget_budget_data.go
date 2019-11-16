// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html

package budgets

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var budgetBudgetDataProperties map[string]string = map[string]string{
	"budget_limit": "BudgetLimit",
	"time_period": "TimePeriod",
	"time_unit": "TimeUnit",
	"planned_budget_limits": "PlannedBudgetLimits",
	"cost_filters": "CostFilters",
	"budget_name": "BudgetName",
	"cost_types": "CostTypes",
	"budget_type": "BudgetType",
}

func propertyBudgetBudgetData(extras...string) *schema.Resource {
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
			"budget_limit": {
				Type: schema.TypeSet,
				Elem: propertyBudgetSpend(),
				Optional: true,
				MaxItems: 1,
			},
			"time_period": {
				Type: schema.TypeSet,
				Elem: propertyBudgetTimePeriod(),
				Optional: true,
				MaxItems: 1,
			},
			"time_unit": {
				Type: schema.TypeString,
				Required: true,
			},
			"planned_budget_limits": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"cost_filters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"budget_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cost_types": {
				Type: schema.TypeSet,
				Elem: propertyBudgetCostTypes(),
				Optional: true,
				MaxItems: 1,
			},
			"budget_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
