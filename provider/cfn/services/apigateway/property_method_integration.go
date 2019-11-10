// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html

package apigateway

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMethodIntegration(extras...string) *schema.Resource {
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
			"cache_key_parameters": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"cache_namespace": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connection_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connection_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"content_handling": {
				Type: schema.TypeString,
				Optional: true,
			},
			"credentials": {
				Type: schema.TypeString,
				Optional: true,
			},
			"integration_http_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"integration_responses": {
				Type: schema.TypeSet,
				Elem: propertyMethodIntegrationResponse(),
				Optional: true,
			},
			"passthrough_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"request_templates": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"timeout_in_millis": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}