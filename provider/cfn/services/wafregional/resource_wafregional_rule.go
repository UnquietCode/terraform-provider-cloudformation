// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalRuleExists,
		Read: resourceWAFRegionalRuleRead,
		Create: resourceWAFRegionalRuleCreate,
		Update: resourceWAFRegionalRuleUpdate,
		Delete: resourceWAFRegionalRuleDelete,
		CustomizeDiff: resourceWAFRegionalRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"predicates": {
				Type: schema.TypeList,
				Elem: propertyRulePredicate(),
				Optional: true,
			},
			"name": {
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

func resourceWAFRegionalRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::Rule", ResourceWAFRegionalRule(), data, meta)
}

func resourceWAFRegionalRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::Rule", ResourceWAFRegionalRule(), data, meta)
}

func resourceWAFRegionalRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::Rule", ResourceWAFRegionalRule(), data, meta)
}

func resourceWAFRegionalRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::Rule", data, meta)
}

func resourceWAFRegionalRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::WAFRegional::Rule", data, meta)
}

