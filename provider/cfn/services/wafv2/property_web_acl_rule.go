// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html

package wafv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWebACLRule(extras...string) *schema.Resource {
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
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"statement": {
				Type: schema.TypeList,
				Elem: propertyWebACLStatementOne(),
				Optional: true,
				MaxItems: 1,
			},
			"action": {
				Type: schema.TypeList,
				Elem: propertyWebACLRuleAction(),
				Optional: true,
				MaxItems: 1,
			},
			"override_action": {
				Type: schema.TypeList,
				Elem: propertyWebACLOverrideAction(),
				Optional: true,
				MaxItems: 1,
			},
			"visibility_config": {
				Type: schema.TypeList,
				Elem: propertyWebACLVisibilityConfig(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
