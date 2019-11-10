// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package budgets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBudget() *schema.Resource {
	return &schema.Resource{
		Create: resourceBudgetCreate,
		Read:   resourceBudgetRead,
		Update: resourceBudgetUpdate,
		Delete: resourceBudgetDelete,

		Schema: map[string]*schema.Schema{
			"notifications_with_subscribers": {
				Type: schema.TypeList,
				Elem: propertyBudgetNotificationWithSubscribers(),
				Required: false,
				ForceNew: true,
			},
			"budget": {
				Type: schema.TypeList,
				Elem: propertyBudgetBudgetData(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceBudgetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Budgets::Budget", data, meta)
}

func resourceBudgetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Budgets::Budget", data, meta)
}

func resourceBudgetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Budgets::Budget", data, meta)
}

func resourceBudgetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Budgets::Budget", data, meta)
}