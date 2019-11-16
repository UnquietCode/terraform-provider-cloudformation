// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalRateBasedRuleType string = "AWS::WAFRegional::RateBasedRule"

func ResourceWAFRegionalRateBasedRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalRateBasedRuleExists,
		Read: resourceWAFRegionalRateBasedRuleRead,
		Create: resourceWAFRegionalRateBasedRuleCreate,
		Update: resourceWAFRegionalRateBasedRuleUpdate,
		Delete: resourceWAFRegionalRateBasedRuleDelete,
		CustomizeDiff: resourceWAFRegionalRateBasedRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rate_limit": {
				Type: schema.TypeInt,
				Required: true,
			},
			"match_predicates": {
				Type: schema.TypeList,
				Elem: propertyRateBasedRulePredicate(),
				Optional: true,
			},
			"rate_key": {
				Type: schema.TypeString,
				Required: true,
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

func resourceWAFRegionalRateBasedRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalRateBasedRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalRateBasedRuleType, ResourceWAFRegionalRateBasedRule(), data, meta)
}

func resourceWAFRegionalRateBasedRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalRateBasedRuleType, ResourceWAFRegionalRateBasedRule(), data, meta)
}

func resourceWAFRegionalRateBasedRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalRateBasedRuleType, ResourceWAFRegionalRateBasedRule(), data, meta)
}

func resourceWAFRegionalRateBasedRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalRateBasedRuleType, data, meta)
}

func resourceWAFRegionalRateBasedRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalRateBasedRuleType, data, meta)
}
