// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingPlansScalingPlanType string = "AWS::AutoScalingPlans::ScalingPlan"

var autoScalingPlansScalingPlanProperties map[string]string = map[string]string{
	"application_source": "ApplicationSource",
	"scaling_instructions": "ScalingInstructions",
}

func ResourceAutoScalingPlansScalingPlan() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAutoScalingPlansScalingPlanExists,
		Read: resourceAutoScalingPlansScalingPlanRead,
		Create: resourceAutoScalingPlansScalingPlanCreate,
		Update: resourceAutoScalingPlansScalingPlanUpdate,
		Delete: resourceAutoScalingPlansScalingPlanDelete,
		CustomizeDiff: resourceAutoScalingPlansScalingPlanCustomizeDiff,
		
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingPlansScalingPlanExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAutoScalingPlansScalingPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingPlansScalingPlanType, ResourceAutoScalingPlansScalingPlan(), data, meta)
}

func resourceAutoScalingPlansScalingPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(autoScalingPlansScalingPlanType, ResourceAutoScalingPlansScalingPlan(), data, autoScalingPlansScalingPlanProperties, meta)
}

func resourceAutoScalingPlansScalingPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(autoScalingPlansScalingPlanType, ResourceAutoScalingPlansScalingPlan(), data, autoScalingPlansScalingPlanProperties, meta)
}

func resourceAutoScalingPlansScalingPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(autoScalingPlansScalingPlanType, data, meta)
}

func resourceAutoScalingPlansScalingPlanCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(autoScalingPlansScalingPlanType, data, meta)
}
