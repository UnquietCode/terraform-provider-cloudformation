// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterCloudWatchAlarmDefinition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparison_operator": {
				Type: schema.TypeString,
				Required: true,
			},
			"dimensions": {
				Type: schema.TypeSet,
				Elem: propertyClusterMetricDimension(),
				Required: false,
			},
			"evaluation_periods": {
				Type: schema.TypeInt,
				Required: false,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type: schema.TypeString,
				Required: false,
			},
			"period": {
				Type: schema.TypeInt,
				Required: true,
			},
			"statistic": {
				Type: schema.TypeString,
				Required: false,
			},
			"threshold": {
				Type: schema.TypeFloat,
				Required: true,
			},
			"unit": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}