// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySpotFleetSpotFleetRequestConfigData(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allocation_strategy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
				Required: false,
			},
			"iam_fleet_role": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_interruption_behavior": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"launch_specifications": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetSpotFleetLaunchSpecification(),
				Required: false,
				ForceNew: true,
			},
			"launch_template_configs": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetLaunchTemplateConfig(),
				Required: false,
				ForceNew: true,
			},
			"load_balancers_config": {
				Type: schema.TypeList,
				Elem: propertySpotFleetLoadBalancersConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"spot_price": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"target_capacity": {
				Type: schema.TypeInt,
				Required: true,
			},
			"terminate_instances_with_expiration": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"valid_from": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"valid_until": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}