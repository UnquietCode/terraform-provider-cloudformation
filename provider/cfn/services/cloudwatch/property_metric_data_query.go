// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMetricDataQuery() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"label": {
				Type: schema.TypeString,
				Required: false,
			},
			"metric_stat": {
				Type: schema.TypeList,
				Elem: propertyMetricStat(),
				Required: false,
				MaxItems: 1,
			},
			"return_data": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}