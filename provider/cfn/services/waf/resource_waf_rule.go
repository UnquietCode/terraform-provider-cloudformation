// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRuleType string = "AWS::WAF::Rule"

func ResourceWAFRule() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRuleRead,
		Create: resourceWAFRuleCreate,
		Update: resourceWAFRuleUpdate,
		Delete: resourceWAFRuleDelete,
		CustomizeDiff: resourceWAFRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"predicates": {
				Type: schema.TypeSet,
				Elem: propertyRulePredicate(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceWAFRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRuleType, ResourceWAFRule(), data, meta)
}

func resourceWAFRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRuleType, ResourceWAFRule(), data, meta)
}

func resourceWAFRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRuleType, ResourceWAFRule(), data, meta)
}

func resourceWAFRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRuleType, data, meta)
}

func resourceWAFRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRuleType, data, meta)
}
