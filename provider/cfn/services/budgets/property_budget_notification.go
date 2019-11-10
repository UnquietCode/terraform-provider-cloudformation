// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package budgets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBudgetNotification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparison_operator": {
				Type: schema.TypeString,
				Required: true,
			},
			"notification_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"threshold": {
				Type: schema.TypeFloat,
				Required: true,
			},
			"threshold_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}