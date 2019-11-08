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

func propertyTargetTrackingConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"customized_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyCustomizedMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"disable_scale_in": {
				Type: schema.TypeBool,
				Required: false,
			},
			"predefined_metric_specification": {
				Type: schema.TypeList,
				Elem: propertyPredefinedMetricSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"target_value": {
				Type: schema.TypeFloat,
				Required: true,
			},
		},
	}
}