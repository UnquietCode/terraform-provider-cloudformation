// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftFleetType string = "AWS::GameLift::Fleet"

var gameLiftFleetProperties map[string]string = map[string]string{
	"build_id": "BuildId",
	"description": "Description",
	"desired_ec2_instances": "DesiredEC2Instances",
	"ec2_inbound_permissions": "EC2InboundPermissions",
	"ec2_instance_type": "EC2InstanceType",
	"log_paths": "LogPaths",
	"max_size": "MaxSize",
	"min_size": "MinSize",
	"name": "Name",
	"server_launch_parameters": "ServerLaunchParameters",
	"server_launch_path": "ServerLaunchPath",
}

func ResourceGameLiftFleet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGameLiftFleetExists,
		Read: resourceGameLiftFleetRead,
		Create: resourceGameLiftFleetCreate,
		Update: resourceGameLiftFleetUpdate,
		Delete: resourceGameLiftFleetDelete,
		CustomizeDiff: resourceGameLiftFleetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"build_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"desired_ec2_instances": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ec2_inbound_permissions": {
				Type: schema.TypeSet,
				Elem: propertyFleetIpPermission(),
				Optional: true,
			},
			"ec2_instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_paths": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"max_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"min_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"server_launch_parameters": {
				Type: schema.TypeString,
				Optional: true,
			},
			"server_launch_path": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGameLiftFleetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGameLiftFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftFleetType, ResourceGameLiftFleet(), data, meta)
}

func resourceGameLiftFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(gameLiftFleetType, ResourceGameLiftFleet(), data, gameLiftFleetProperties, meta)
}

func resourceGameLiftFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(gameLiftFleetType, ResourceGameLiftFleet(), data, gameLiftFleetProperties, meta)
}

func resourceGameLiftFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(gameLiftFleetType, data, meta)
}

func resourceGameLiftFleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(gameLiftFleetType, data, meta)
}
