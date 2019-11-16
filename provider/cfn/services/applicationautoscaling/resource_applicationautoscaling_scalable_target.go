// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const applicationAutoScalingScalableTargetType string = "AWS::ApplicationAutoScaling::ScalableTarget"

var applicationAutoScalingScalableTargetProperties map[string]string = map[string]string{
	"max_capacity": "MaxCapacity",
	"min_capacity": "MinCapacity",
	"resource_id": "ResourceId",
	"role_arn": "RoleARN",
	"scalable_dimension": "ScalableDimension",
	"scheduled_actions": "ScheduledActions",
	"service_namespace": "ServiceNamespace",
	"suspended_state": "SuspendedState",
}

func ResourceApplicationAutoScalingScalableTarget() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApplicationAutoScalingScalableTargetExists,
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
				Type: schema.TypeList,
				Elem: propertyScalableTargetSuspendedState(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApplicationAutoScalingScalableTargetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApplicationAutoScalingScalableTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, applicationAutoScalingScalableTargetProperties, meta)
}

func resourceApplicationAutoScalingScalableTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(applicationAutoScalingScalableTargetType, ResourceApplicationAutoScalingScalableTarget(), data, applicationAutoScalingScalableTargetProperties, meta)
}

func resourceApplicationAutoScalingScalableTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(applicationAutoScalingScalableTargetType, data, meta)
}

func resourceApplicationAutoScalingScalableTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(applicationAutoScalingScalableTargetType, data, meta)
}
