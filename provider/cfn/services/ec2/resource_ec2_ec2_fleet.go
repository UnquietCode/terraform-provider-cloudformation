// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceEC2EC2FleetExists,
		Read: resourceEC2EC2FleetRead,
		Create: resourceEC2EC2FleetCreate,
		Update: resourceEC2EC2FleetUpdate,
		Delete: resourceEC2EC2FleetDelete,
		CustomizeDiff: resourceEC2EC2FleetCustomizeDiff,
		
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
				Optional: true,
				MaxItems: 1,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetTagSpecification(),
				Optional: true,
			},
			"spot_options": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetSpotOptionsRequest(),
				Optional: true,
				MaxItems: 1,
			},
			"valid_from": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"launch_template_configs": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateConfigRequest(),
				Required: true,
			},
			"terminate_instances_with_expiration": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"valid_until": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2EC2FleetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2EC2FleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EC2Fleet", ResourceEC2EC2Fleet(), data, meta)
}

func resourceEC2EC2FleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EC2Fleet", ResourceEC2EC2Fleet(), data, meta)
}

func resourceEC2EC2FleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EC2Fleet", ResourceEC2EC2Fleet(), data, meta)
}

func resourceEC2EC2FleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EC2Fleet", data, meta)
}

func resourceEC2EC2FleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

