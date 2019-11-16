// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html

package opsworks

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var layerAutoScalingThresholdsProperties map[string]string = map[string]string{
	"cpu_threshold": "CpuThreshold",
	"ignore_metrics_time": "IgnoreMetricsTime",
	"instance_count": "InstanceCount",
	"load_threshold": "LoadThreshold",
	"memory_threshold": "MemoryThreshold",
	"thresholds_wait_time": "ThresholdsWaitTime",
}

func propertyLayerAutoScalingThresholds(extras...string) *schema.Resource {
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
			"cpu_threshold": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"ignore_metrics_time": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"load_threshold": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"memory_threshold": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"thresholds_wait_time": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
