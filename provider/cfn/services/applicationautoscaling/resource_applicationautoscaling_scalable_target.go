// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplicationAutoScalingScalableTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationAutoScalingScalableTargetCreate,
		Read:   resourceApplicationAutoScalingScalableTargetRead,
		Update: resourceApplicationAutoScalingScalableTargetUpdate,
		Delete: resourceApplicationAutoScalingScalableTargetDelete,

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
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"scalable_dimension": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scheduled_actions": {
				Type: schema.TypeSet,
				Elem: propertyScalableTargetScheduledAction(),
				Optional: true,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"suspended_state": {
				Type: schema.TypeList,
				Elem: propertyScalableTargetSuspendedState(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApplicationAutoScalingScalableTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApplicationAutoScaling::ScalableTarget", ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApplicationAutoScaling::ScalableTarget", ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApplicationAutoScaling::ScalableTarget", ResourceApplicationAutoScalingScalableTarget(), data, meta)
}

func resourceApplicationAutoScalingScalableTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApplicationAutoScaling::ScalableTarget", data, meta)
}