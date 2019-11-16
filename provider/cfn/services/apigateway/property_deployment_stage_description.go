// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html

package apigateway

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
)

var deploymentStageDescriptionProperties map[string]string = map[string]string{
	"access_log_setting": "AccessLogSetting",
	"cache_cluster_enabled": "CacheClusterEnabled",
	"cache_cluster_size": "CacheClusterSize",
	"cache_data_encrypted": "CacheDataEncrypted",
	"cache_ttl_in_seconds": "CacheTtlInSeconds",
	"caching_enabled": "CachingEnabled",
	"canary_setting": "CanarySetting",
	"client_certificate_id": "ClientCertificateId",
	"data_trace_enabled": "DataTraceEnabled",
	"description": "Description",
	"documentation_version": "DocumentationVersion",
	"logging_level": "LoggingLevel",
	"method_settings": "MethodSettings",
	"metrics_enabled": "MetricsEnabled",
	"tags": "Tags",
	"throttling_burst_limit": "ThrottlingBurstLimit",
	"throttling_rate_limit": "ThrottlingRateLimit",
	"tracing_enabled": "TracingEnabled",
	"variables": "Variables",
}

func propertyDeploymentStageDescription(extras...string) *schema.Resource {
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
			"access_log_setting": {
				Type: schema.TypeList,
				Elem: propertyDeploymentAccessLogSetting(),
				Optional: true,
				MaxItems: 1,
			},
			"cache_cluster_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cache_cluster_size": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cache_data_encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cache_ttl_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"caching_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"canary_setting": {
				Type: schema.TypeList,
				Elem: propertyDeploymentCanarySetting(),
				Optional: true,
				MaxItems: 1,
			},
			"client_certificate_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logging_level": {
				Type: schema.TypeString,
				Optional: true,
			},
			"method_settings": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentMethodSetting(),
				Optional: true,
			},
			"metrics_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"throttling_burst_limit": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"throttling_rate_limit": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"tracing_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"variables": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}
