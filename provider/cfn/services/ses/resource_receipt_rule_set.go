// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceReceiptRuleSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceReceiptRuleSetCreate,
		Read:   resourceReceiptRuleSetRead,
		Update: resourceReceiptRuleSetUpdate,
		Delete: resourceReceiptRuleSetDelete,

		Schema: map[string]*schema.Schema{
			"rule_set_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceReceiptRuleSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceReceiptRuleSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceReceiptRuleSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceReceiptRuleSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptRuleSet", data, meta)
}