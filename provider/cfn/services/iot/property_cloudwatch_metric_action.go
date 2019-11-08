// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCloudwatchMetricAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_timestamp": {
				Type: schema.TypeString,
				Required: false,
			},
			"metric_unit": {
				Type: schema.TypeString,
				Required: true,
			},
			"metric_value": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}