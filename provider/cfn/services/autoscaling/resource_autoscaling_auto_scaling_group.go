// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html

package autoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingAutoScalingGroupType string = "AWS::AutoScaling::AutoScalingGroup"

func ResourceAutoScalingAutoScalingGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceAutoScalingAutoScalingGroupRead,
		Create: resourceAutoScalingAutoScalingGroupCreate,
		Update: resourceAutoScalingAutoScalingGroupUpdate,
		Delete: resourceAutoScalingAutoScalingGroupDelete,
		CustomizeDiff: resourceAutoScalingAutoScalingGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_scaling_group_name": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"launch_configuration_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_template": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAutoScalingAutoScalingGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingAutoScalingGroupType, ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(autoScalingAutoScalingGroupType, ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(autoScalingAutoScalingGroupType, ResourceAutoScalingAutoScalingGroup(), data, meta)
}

func resourceAutoScalingAutoScalingGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(autoScalingAutoScalingGroupType, data, meta)
}

func resourceAutoScalingAutoScalingGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(autoScalingAutoScalingGroupType, data, meta)
}
