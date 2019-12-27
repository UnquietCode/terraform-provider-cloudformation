// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html

package wafv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWebACLStatementOne(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
	    if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
	        count = i
	    }
	}
	
	if count >= 5 {
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"byte_match_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLByteMatchStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"sqli_match_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLSqliMatchStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"xss_match_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLXssMatchStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"size_constraint_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLSizeConstraintStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"geo_match_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLGeoMatchStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"rule_group_reference_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLRuleGroupReferenceStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"ip_set_reference_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLIPSetReferenceStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"regex_pattern_set_reference_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLRegexPatternSetReferenceStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"managed_rule_group_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLManagedRuleGroupStatement(),
				Optional: true,
				MaxItems: 1,
			},
			"rate_based_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLRateBasedStatementOne(),
				Optional: true,
				MaxItems: 1,
			},
			"and_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLAndStatementOne(),
				Optional: true,
				MaxItems: 1,
			},
			"or_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLOrStatementOne(),
				Optional: true,
				MaxItems: 1,
			},
			"not_statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLNotStatementOne(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
