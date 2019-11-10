// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field": {
				Type: schema.TypeString,
				Required: false,
			},
			"host_header_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHostHeaderConfig(),
				Required: false,
				MaxItems: 1,
			},
			"http_header_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHttpHeaderConfig(),
				Required: false,
				MaxItems: 1,
			},
			"http_request_method_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleHttpRequestMethodConfig(),
				Required: false,
				MaxItems: 1,
			},
			"path_pattern_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRulePathPatternConfig(),
				Required: false,
				MaxItems: 1,
			},
			"query_string_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleQueryStringConfig(),
				Required: false,
				MaxItems: 1,
			},
			"source_ip_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleSourceIpConfig(),
				Required: false,
				MaxItems: 1,
			},
			"values": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
		},
	}
}