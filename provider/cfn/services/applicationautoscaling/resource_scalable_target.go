// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package applicationautoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceScalableTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceScalableTargetCreate,
		Read:   resourceScalableTargetRead,
		Update: resourceScalableTargetUpdate,
		Delete: resourceScalableTargetDelete,

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
				Required: false,
			},
			"service_namespace": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"suspended_state": {
				Type: schema.TypeList,
				Elem: propertyScalableTargetSuspendedState(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceScalableTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApplicationAutoScaling::ScalableTarget", data, meta)
}

func resourceScalableTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApplicationAutoScaling::ScalableTarget", data, meta)
}

func resourceScalableTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApplicationAutoScaling::ScalableTarget", data, meta)
}

func resourceScalableTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApplicationAutoScaling::ScalableTarget", data, meta)
}