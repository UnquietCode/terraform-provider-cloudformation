// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingThresholds() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_threshold": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"ignore_metrics_time": {
				Type: schema.TypeInt,
				Required: false,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"load_threshold": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"memory_threshold": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"thresholds_wait_time": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}