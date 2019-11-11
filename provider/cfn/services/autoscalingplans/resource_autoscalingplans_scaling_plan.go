// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingPlansScalingPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutoScalingPlansScalingPlanCreate,
		Read:   resourceAutoScalingPlansScalingPlanRead,
		Update: resourceAutoScalingPlansScalingPlanUpdate,
		Delete: resourceAutoScalingPlansScalingPlanDelete,

		Schema: map[string]*schema.Schema{
			"scaling_plan_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"scaling_plan_version": {
				Type: schema.TypeString,
				Computed: true,
			},
			"application_source": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanApplicationSource(),
				Required: true,
				MaxItems: 1,
			},
			"scaling_instructions": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanScalingInstruction(),
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingPlansScalingPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScalingPlans::ScalingPlan", ResourceAutoScalingPlansScalingPlan(), data, meta)
}

func resourceAutoScalingPlansScalingPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScalingPlans::ScalingPlan", ResourceAutoScalingPlansScalingPlan(), data, meta)
}

func resourceAutoScalingPlansScalingPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScalingPlans::ScalingPlan", ResourceAutoScalingPlansScalingPlan(), data, meta)
}

func resourceAutoScalingPlansScalingPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScalingPlans::ScalingPlan", data, meta)
}