// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html

package wafregional

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalRateBasedRuleType string = "AWS::WAFRegional::RateBasedRule"

func DatasourceWAFRegionalRateBasedRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWAFRegionalRateBasedRuleRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceWAFRegionalRateBasedRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalRateBasedRuleType, DatasourceWAFRegionalRateBasedRule(), data, meta)
}
