// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCustomizedMetricSpecification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dimensions": {
				Type: schema.TypeSet,
				Elem: propertyMetricDimension(),
				Required: false,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"statistic": {
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