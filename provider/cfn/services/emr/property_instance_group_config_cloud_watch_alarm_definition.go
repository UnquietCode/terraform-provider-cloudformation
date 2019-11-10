// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceGroupConfigCloudWatchAlarmDefinition(extras...string) *schema.Resource {
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
			"comparison_operator": {
				Type: schema.TypeString,
				Required: true,
			},
			"dimensions": {
				Type: schema.TypeSet,
				Elem: propertyInstanceGroupConfigMetricDimension(),
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