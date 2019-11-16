// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html

package budgets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const budgetsBudgetType string = "AWS::Budgets::Budget"

func ResourceBudgetsBudget() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBudgetsBudgetExists,
		Read: resourceBudgetsBudgetRead,
		Create: resourceBudgetsBudgetCreate,
		Update: resourceBudgetsBudgetUpdate,
		Delete: resourceBudgetsBudgetDelete,
		CustomizeDiff: resourceBudgetsBudgetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"notifications_with_subscribers": {
				Type: schema.TypeList,
				Elem: propertyBudgetNotificationWithSubscribers(),
				Optional: true,
			},
			"budget": {
				Type: schema.TypeSet,
				Elem: propertyBudgetBudgetData(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBudgetsBudgetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceBudgetsBudgetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(budgetsBudgetType, ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(budgetsBudgetType, ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(budgetsBudgetType, ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(budgetsBudgetType, data, meta)
}

func resourceBudgetsBudgetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(budgetsBudgetType, data, meta)
}
