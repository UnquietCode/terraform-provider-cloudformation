// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html

package applicationautoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const applicationAutoScalingScalableTargetType string = "AWS::ApplicationAutoScaling::ScalableTarget"

func ResourceApplicationAutoScalingScalableTarget() *schema.Resource {
	return &schema.Resource{
		Read: resourceApplicationAutoScalingScalableTargetRead,
		Create: resourceApplicationAutoScalingScalableTargetCreate,
		Update: resourceApplicationAutoScalingScalableTargetUpdate,
		Delete: resourceApplicationAutoScalingScalableTargetDelete,
		CustomizeDiff: resourceApplicationAutoScalingScalableTargetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"max_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"min_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Required: true,
			},
			"scheduled_actions": {
				Type: schema.TypeSet,
				Elem: propertyScalableTargetScheduledAction(),
				Optional: true,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"suspended_state": {
				Type: schema.TypeSet,
				Elem: propertyScalableTargetSuspendedState(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceApplicationAutoScalingScalableTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(applicationAutoScalingScalableTargetType, data, meta)
}

func resourceApplicationAutoScalingScalableTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(applicationAutoScalingScalableTargetType, data, meta)
}
