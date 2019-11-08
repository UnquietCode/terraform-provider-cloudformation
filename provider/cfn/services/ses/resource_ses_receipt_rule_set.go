// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESReceiptRuleSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSESReceiptRuleSetCreate,
		Read:   resourceSESReceiptRuleSetRead,
		Update: resourceSESReceiptRuleSetUpdate,
		Delete: resourceSESReceiptRuleSetDelete,

		Schema: map[string]*schema.Schema{
			"rule_set_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSESReceiptRuleSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceSESReceiptRuleSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceSESReceiptRuleSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceSESReceiptRuleSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptRuleSet", data, meta)
}