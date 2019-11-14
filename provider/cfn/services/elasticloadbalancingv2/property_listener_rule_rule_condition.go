// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html

package elasticloadbalancingv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyListenerRuleRuleCondition(extras...string) *schema.Resource {
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
			"field": {
				Type: schema.TypeString,
				Optional: true,
			},
			"host_header_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHostHeaderConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"http_header_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHttpHeaderConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"http_request_method_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHttpRequestMethodConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"path_pattern_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRulePathPatternConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"query_string_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleQueryStringConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"source_ip_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleSourceIpConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"values": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
		},
	}
}

