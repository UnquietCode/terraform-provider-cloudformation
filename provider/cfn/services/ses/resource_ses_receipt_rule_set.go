// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptruleset.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESReceiptRuleSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSESReceiptRuleSetExists,
		Read:   resourceSESReceiptRuleSetRead,
		Create: resourceSESReceiptRuleSetCreate,
		Update: resourceSESReceiptRuleSetUpdate,
		Delete: resourceSESReceiptRuleSetDelete,
		CustomizeDiff: resourceSESReceiptRuleSetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"rule_set_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSESReceiptRuleSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSESReceiptRuleSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ReceiptRuleSet", ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ReceiptRuleSet", ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ReceiptRuleSet", ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ReceiptRuleSet", data, meta)
}

func resourceSESReceiptRuleSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}