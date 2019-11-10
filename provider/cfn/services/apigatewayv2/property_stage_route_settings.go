// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStageRouteSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging_level": {
				Type: schema.TypeString,
				Required: false,
			},
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"throttling_burst_limit": {
				Type: schema.TypeInt,
				Required: false,
			},
			"detailed_metrics_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"throttling_rate_limit": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}