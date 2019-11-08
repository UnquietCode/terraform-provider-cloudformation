// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRuleCondition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field": {
				Type: schema.TypeString,
				Required: false,
			},
			"host_header_config": {
				Type: schema.TypeList,
				Elem: propertyHostHeaderConfig(),
				Required: false,
				MaxItems: 1,
			},
			"http_header_config": {
				Type: schema.TypeList,
				Elem: propertyHttpHeaderConfig(),
				Required: false,
				MaxItems: 1,
			},
			"http_request_method_config": {
				Type: schema.TypeList,
				Elem: propertyHttpRequestMethodConfig(),
				Required: false,
				MaxItems: 1,
			},
			"path_pattern_config": {
				Type: schema.TypeList,
				Elem: propertyPathPatternConfig(),
				Required: false,
				MaxItems: 1,
			},
			"query_string_config": {
				Type: schema.TypeList,
				Elem: propertyQueryStringConfig(),
				Required: false,
				MaxItems: 1,
			},
			"source_ip_config": {
				Type: schema.TypeList,
				Elem: propertySourceIpConfig(),
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