// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceScalingPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceScalingPlanCreate,
		Read:   resourceScalingPlanRead,
		Update: resourceScalingPlanUpdate,
		Delete: resourceScalingPlanDelete,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceScalingPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScalingPlans::ScalingPlan", data, meta)
}

func resourceScalingPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScalingPlans::ScalingPlan", data, meta)
}

func resourceScalingPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScalingPlans::ScalingPlan", data, meta)
}

func resourceScalingPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScalingPlans::ScalingPlan", data, meta)
}