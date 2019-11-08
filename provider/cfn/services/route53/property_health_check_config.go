// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyHealthCheckConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alarm_identifier": {
				Type: schema.TypeList,
				Elem: propertyAlarmIdentifier(),
				Required: false,
				MaxItems: 1,
			},
			"child_health_checks": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"enable_sni": {
				Type: schema.TypeBool,
				Required: false,
			},
			"failure_threshold": {
				Type: schema.TypeInt,
				Required: false,
			},
			"fully_qualified_domain_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_threshold": {
				Type: schema.TypeInt,
				Required: false,
			},
			"ip_address": {
				Type: schema.TypeString,
				Required: false,
			},
			"insufficient_data_health_status": {
				Type: schema.TypeString,
				Required: false,
			},
			"inverted": {
				Type: schema.TypeBool,
				Required: false,
			},
			"measure_latency": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"regions": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"request_interval": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"resource_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"search_string": {
				Type: schema.TypeString,
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}