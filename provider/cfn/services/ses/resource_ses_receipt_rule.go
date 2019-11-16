// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sESReceiptRuleType string = "AWS::SES::ReceiptRule"

var sESReceiptRuleProperties map[string]string = map[string]string{
	"after": "After",
	"rule": "Rule",
	"rule_set_name": "RuleSetName",
}

func ResourceSESReceiptRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSESReceiptRuleExists,
		Read: resourceSESReceiptRuleRead,
		Create: resourceSESReceiptRuleCreate,
		Update: resourceSESReceiptRuleUpdate,
		Delete: resourceSESReceiptRuleDelete,
		CustomizeDiff: resourceSESReceiptRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"after": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rule": {
				Type: schema.TypeSet,
				Elem: propertyReceiptRuleRule(),
				Required: true,
				MaxItems: 1,
			},
			"rule_set_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSESReceiptRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSESReceiptRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sESReceiptRuleType, ResourceSESReceiptRule(), data, meta)
}

func resourceSESReceiptRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sESReceiptRuleType, ResourceSESReceiptRule(), data, sESReceiptRuleProperties, meta)
}

func resourceSESReceiptRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sESReceiptRuleType, ResourceSESReceiptRule(), data, sESReceiptRuleProperties, meta)
}

func resourceSESReceiptRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sESReceiptRuleType, data, meta)
}

func resourceSESReceiptRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sESReceiptRuleType, data, meta)
}
