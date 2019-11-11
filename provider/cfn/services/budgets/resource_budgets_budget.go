// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html

package budgets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBudgetsBudget() *schema.Resource {
	return &schema.Resource{
		Create: resourceBudgetsBudgetCreate,
		Read:   resourceBudgetsBudgetRead,
		Update: resourceBudgetsBudgetUpdate,
		Delete: resourceBudgetsBudgetDelete,

		Schema: map[string]*schema.Schema{
			"notifications_with_subscribers": {
				Type: schema.TypeList,
				Elem: propertyBudgetNotificationWithSubscribers(),
				Optional: true,
				ForceNew: true,
			},
			"budget": {
				Type: schema.TypeList,
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

func resourceBudgetsBudgetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Budgets::Budget", ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Budgets::Budget", ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Budgets::Budget", ResourceBudgetsBudget(), data, meta)
}

func resourceBudgetsBudgetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Budgets::Budget", data, meta)
}