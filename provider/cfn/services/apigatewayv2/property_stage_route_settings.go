// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-routesettings.html

package apigatewayv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStageRouteSettings(extras...string) *schema.Resource {
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
			"logging_level": {
				Type: schema.TypeString,
				Optional: true,
			},
			"data_trace_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"throttling_burst_limit": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"detailed_metrics_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"throttling_rate_limit": {
				Type: schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
