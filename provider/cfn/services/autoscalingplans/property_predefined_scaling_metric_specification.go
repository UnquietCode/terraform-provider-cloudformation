// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPredefinedScalingMetricSpecification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_label": {
				Type: schema.TypeString,
				Required: false,
			},
			"predefined_scaling_metric_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}