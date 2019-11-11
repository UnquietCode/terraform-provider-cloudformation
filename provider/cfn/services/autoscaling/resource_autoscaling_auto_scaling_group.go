// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingAutoScalingGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutoScalingAutoScalingGroupCreate,
		Read:   resourceAutoScalingAutoScalingGroupRead,
		Update: resourceAutoScalingAutoScalingGroupUpdate,
		Delete: resourceAutoScalingAutoScalingGroupDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"cooldown": {
				Type: schema.TypeString,
				Optional: true,
			},
			"desired_capacity": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_grace_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"health_check_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"launch_configuration_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupLaunchTemplateSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"lifecycle_hook_specification_list": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupLifecycleHookSpecification(),
				Optional: true,
			},
			"load_balancer_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"max_size": {
				Type: schema.TypeString,
				Required: true,
			},
			"metrics_collection": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupMetricsCollection(),
				Optional: true,
			},
			"min_size": {
				Type: schema.TypeString,
				Required: true,
			},
			"mixed_instances_policy": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupMixedInstancesPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"notification_configurations": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupNotificationConfiguration(),
				Optional: true,
			},
			"placement_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_linked_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyAutoScalingGroupTagProperty(),
				Optional: true,
			},
			"target_group_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"termination_policies": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"vpc_zone_identifier": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingAutoScalingGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::AutoScalingGroup", ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::AutoScalingGroup", ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::AutoScalingGroup", ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::AutoScalingGroup", data, meta)
}