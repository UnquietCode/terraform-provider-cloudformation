// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptruleset.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sESReceiptRuleSetType string = "AWS::SES::ReceiptRuleSet"

func ResourceSESReceiptRuleSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceSESReceiptRuleSetRead,
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

func resourceSESReceiptRuleSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sESReceiptRuleSetType, ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sESReceiptRuleSetType, ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sESReceiptRuleSetType, ResourceSESReceiptRuleSet(), data, meta)
}

func resourceSESReceiptRuleSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sESReceiptRuleSetType, data, meta)
}

func resourceSESReceiptRuleSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sESReceiptRuleSetType, data, meta)
}
