// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

func propertyDeploymentStageDescription() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_log_setting": {
				Type: schema.TypeList,
				Elem: propertyDeploymentAccessLogSetting(),
				Required: false,
				MaxItems: 1,
			},
			"cache_cluster_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"cache_cluster_size": {
				Type: schema.TypeString,
				Required: false,
			},
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
			"canary_setting": {
				Type: schema.TypeList,
				Elem: propertyDeploymentCanarySetting(),
				Required: false,
				MaxItems: 1,
			},
			"client_certificate_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"logging_level": {
				Type: schema.TypeString,
				Required: false,
			},
			"method_settings": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentMethodSetting(),
				Required: false,
			},
			"metrics_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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
			"tracing_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"variables": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}