// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var applicationFlinkApplicationConfigurationProperties map[string]string = map[string]string{
	"checkpoint_configuration": "CheckpointConfiguration",
	"parallelism_configuration": "ParallelismConfiguration",
	"monitoring_configuration": "MonitoringConfiguration",
}

func propertyApplicationFlinkApplicationConfiguration(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checkpoint_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationCheckpointConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"parallelism_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationParallelismConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"monitoring_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationMonitoringConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
