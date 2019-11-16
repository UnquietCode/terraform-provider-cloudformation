// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html

package autoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var scalingPolicyTargetTrackingConfigurationProperties map[string]string = map[string]string{
	"customized_metric_specification": "CustomizedMetricSpecification",
	"disable_scale_in": "DisableScaleIn",
	"predefined_metric_specification": "PredefinedMetricSpecification",
	"target_value": "TargetValue",
}

func propertyScalingPolicyTargetTrackingConfiguration(extras...string) *schema.Resource {
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
			"customized_metric_specification": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyCustomizedMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"disable_scale_in": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"predefined_metric_specification": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyPredefinedMetricSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"target_value": {
				Type: schema.TypeFloat,
				Required: true,
			},
		},
	}
}
