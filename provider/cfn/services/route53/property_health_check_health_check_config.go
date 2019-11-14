// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html

package route53

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyHealthCheckHealthCheckConfig(extras...string) *schema.Resource {
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
			"alarm_identifier": {
				Type: schema.TypeList,
				Elem: propertyHealthCheckAlarmIdentifier(),
				Optional: true,
				MaxItems: 1,
			},
			"child_health_checks": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"enable_sni": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"failure_threshold": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"fully_qualified_domain_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_threshold": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ip_address": {
				Type: schema.TypeString,
				Optional: true,
			},
			"insufficient_data_health_status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"inverted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"measure_latency": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"regions": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"request_interval": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"resource_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"search_string": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
