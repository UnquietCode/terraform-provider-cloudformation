// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutoScalingGroupCreate,
		Read:   resourceAutoScalingGroupRead,
		Update: resourceAutoScalingGroupUpdate,
		Delete: resourceAutoScalingGroupDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"cooldown": {
				Type: schema.TypeString,
				Required: false,
			},
			"desired_capacity": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_check_grace_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"health_check_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"launch_configuration_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupLaunchTemplateSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"lifecycle_hook_specification_list": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupLifecycleHookSpecification(),
				Required: false,
			},
			"load_balancer_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"max_size": {
				Type: schema.TypeString,
				Required: true,
			},
			"metrics_collection": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupMetricsCollection(),
				Required: false,
			},
			"min_size": {
				Type: schema.TypeString,
				Required: true,
			},
			"mixed_instances_policy": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupMixedInstancesPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"notification_configurations": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupNotificationConfiguration(),
				Required: false,
			},
			"placement_group": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_linked_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupTagProperty(),
				Required: false,
			},
			"target_group_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"termination_policies": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"vpc_zone_identifier": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceAutoScalingGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::AutoScalingGroup", data, meta)
}

func resourceAutoScalingGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::AutoScalingGroup", data, meta)
}

func resourceAutoScalingGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::AutoScalingGroup", data, meta)
}

func resourceAutoScalingGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::AutoScalingGroup", data, meta)
}