// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var spotFleetSpotFleetRequestConfigDataProperties map[string]string = map[string]string{
	"allocation_strategy": "AllocationStrategy",
	"excess_capacity_termination_policy": "ExcessCapacityTerminationPolicy",
	"iam_fleet_role": "IamFleetRole",
	"instance_interruption_behavior": "InstanceInterruptionBehavior",
	"launch_specifications": "LaunchSpecifications",
	"launch_template_configs": "LaunchTemplateConfigs",
	"load_balancers_config": "LoadBalancersConfig",
	"replace_unhealthy_instances": "ReplaceUnhealthyInstances",
	"spot_price": "SpotPrice",
	"target_capacity": "TargetCapacity",
	"terminate_instances_with_expiration": "TerminateInstancesWithExpiration",
	"type": "Type",
	"valid_from": "ValidFrom",
	"valid_until": "ValidUntil",
}

func propertySpotFleetSpotFleetRequestConfigData(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allocation_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"iam_fleet_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_interruption_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_specifications": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetSpotFleetLaunchSpecification(),
				Optional: true,
			},
			"launch_template_configs": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetLaunchTemplateConfig(),
				Optional: true,
			},
			"load_balancers_config": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetLoadBalancersConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"spot_price": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"terminate_instances_with_expiration": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"valid_from": {
				Type: schema.TypeString,
				Optional: true,
			},
			"valid_until": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
