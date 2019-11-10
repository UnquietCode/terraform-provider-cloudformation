// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMetricFilterMetricTransformation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_value": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_value": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}