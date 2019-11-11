// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html

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
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"budget_limit": {
				Type: schema.TypeList,
				Elem: propertyBudgetSpend(),
				Optional: true,
				MaxItems: 1,
			},
			"time_period": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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

