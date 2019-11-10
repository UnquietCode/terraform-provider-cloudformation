// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationFlinkApplicationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checkpoint_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationCheckpointConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"parallelism_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationParallelismConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"monitoring_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationMonitoringConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}