// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_key_parameters": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"cache_namespace": {
				Type: schema.TypeString,
				Required: false,
			},
			"connection_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"connection_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"content_handling": {
				Type: schema.TypeString,
				Required: false,
			},
			"credentials": {
				Type: schema.TypeString,
				Required: false,
			},
			"integration_http_method": {
				Type: schema.TypeString,
				Required: false,
			},
			"integration_responses": {
				Type: schema.TypeSet,
				Elem: propertyMethodIntegrationResponse(),
				Required: false,
			},
			"passthrough_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"request_templates": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"timeout_in_millis": {
				Type: schema.TypeInt,
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: false,
			},
			"uri": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}