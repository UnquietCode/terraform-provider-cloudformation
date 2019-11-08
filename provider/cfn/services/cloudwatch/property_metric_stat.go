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

func propertyMetricStat() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric": {
				Type: schema.TypeList,
				Elem: propertyMetric(),
				Required: true,
				MaxItems: 1,
			},
			"period": {
				Type: schema.TypeInt,
				Required: true,
			},
			"stat": {
				Type: schema.TypeString,
				Required: true,
			},
			"unit": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}