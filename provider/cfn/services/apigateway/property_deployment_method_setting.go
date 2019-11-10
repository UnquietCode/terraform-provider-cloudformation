// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeploymentMethodSetting() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_data_encrypted": {
				Type: schema.TypeBool,
				Required: false,
			},
			"cache_ttl_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"caching_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"http_method": {
				Type: schema.TypeString,
				Required: false,
			},
			"logging_level": {
				Type: schema.TypeString,
				Required: false,
			},
			"metrics_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"resource_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"throttling_burst_limit": {
				Type: schema.TypeInt,
				Required: false,
			},
			"throttling_rate_limit": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}