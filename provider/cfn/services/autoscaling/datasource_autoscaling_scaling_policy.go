// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html

package autoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingScalingPolicyType string = "AWS::AutoScaling::ScalingPolicy"

func DatasourceAutoScalingScalingPolicy() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAutoScalingScalingPolicyRead,
		
		Schema: map[string]*schema.Schema{
			"adjustment_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cooldown": {
				Type: schema.TypeString,
				Optional: true,
			},
			"estimated_instance_warmup": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"metric_aggregation_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"min_adjustment_magnitude": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"policy_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scaling_adjustment": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"step_adjustments": {
				Type: schema.TypeSet,
				Elem: propertyScalingPolicyStepAdjustment(),
				Optional: true,
			},
			"target_tracking_configuration": {
				Type: schema.TypeList,
				Elem: propertyScalingPolicyTargetTrackingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAutoScalingScalingPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingScalingPolicyType, DatasourceAutoScalingScalingPolicy(), data, meta)
}
