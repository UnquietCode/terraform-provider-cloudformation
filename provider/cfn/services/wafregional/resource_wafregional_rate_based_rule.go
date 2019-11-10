// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalRateBasedRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalRateBasedRuleCreate,
		Read:   resourceWAFRegionalRateBasedRuleRead,
		Update: resourceWAFRegionalRateBasedRuleUpdate,
		Delete: resourceWAFRegionalRateBasedRuleDelete,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rate_limit": {
				Type: schema.TypeInt,
				Required: true,
			},
			"match_predicates": {
				Type: schema.TypeList,
				Elem: propertyRateBasedRulePredicate(),
				Required: false,
			},
			"rate_key": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalRateBasedRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::RateBasedRule", data, meta)
}

func resourceWAFRegionalRateBasedRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::RateBasedRule", data, meta)
}

func resourceWAFRegionalRateBasedRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::RateBasedRule", data, meta)
}

func resourceWAFRegionalRateBasedRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::RateBasedRule", data, meta)
}