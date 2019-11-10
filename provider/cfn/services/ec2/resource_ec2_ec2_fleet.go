// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2EC2Fleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2EC2FleetCreate,
		Read:   resourceEC2EC2FleetRead,
		Update: resourceEC2EC2FleetUpdate,
		Delete: resourceEC2EC2FleetDelete,

		Schema: map[string]*schema.Schema{
			"target_capacity_specification": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetTargetCapacitySpecificationRequest(),
				Required: true,
				MaxItems: 1,
			},
			"on_demand_options": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetOnDemandOptionsRequest(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetTagSpecification(),
				Required: false,
				ForceNew: true,
			},
			"spot_options": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetSpotOptionsRequest(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"valid_from": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"launch_template_configs": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateConfigRequest(),
				Required: true,
				ForceNew: true,
			},
			"terminate_instances_with_expiration": {
				Type: schema.TypeBool,
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

func resourceEC2EC2FleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EC2Fleet", data, meta)
}

func resourceEC2EC2FleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EC2Fleet", data, meta)
}

func resourceEC2EC2FleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EC2Fleet", data, meta)
}

func resourceEC2EC2FleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EC2Fleet", data, meta)
}