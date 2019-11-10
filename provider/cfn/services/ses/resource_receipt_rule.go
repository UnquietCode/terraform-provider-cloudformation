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

func ResourceReceiptRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceReceiptRuleCreate,
		Read:   resourceReceiptRuleRead,
		Update: resourceReceiptRuleUpdate,
		Delete: resourceReceiptRuleDelete,

		Schema: map[string]*schema.Schema{
			"after": {
				Type: schema.TypeString,
				Required: false,
			},
			"rule": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleRule(),
				Required: true,
				MaxItems: 1,
			},
			"rule_set_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceReceiptRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptRule", data, meta)
}

func resourceReceiptRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptRule", data, meta)
}

func resourceReceiptRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptRule", data, meta)
}

func resourceReceiptRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptRule", data, meta)
}