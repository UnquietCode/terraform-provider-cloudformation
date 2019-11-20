// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription-methodsetting.html

package apigateway

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeploymentMethodSetting(extras...string) *schema.Resource {
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
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"http_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logging_level": {
				Type: schema.TypeString,
				Optional: true,
			},
			"metrics_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"resource_path": {
				Type: schema.TypeString,
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
		},
	}
}
