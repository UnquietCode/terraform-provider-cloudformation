// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robot.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerRobot() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerRobotExists,
		Read:   resourceRoboMakerRobotRead,
		Create: resourceRoboMakerRobotCreate,
		Update: resourceRoboMakerRobotUpdate,
		Delete: resourceRoboMakerRobotDelete,
		CustomizeDiff: resourceRoboMakerRobotCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"fleet": {
				Type: schema.TypeString,
				Optional: true,
			},
			"architecture": {
				Type: schema.TypeString,
				Required: true,
			},
			"greengrass_group_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
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

func resourceRoboMakerRobotExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerRobotRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::Robot", ResourceRoboMakerRobot(), data, meta)
}

func resourceRoboMakerRobotCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::Robot", ResourceRoboMakerRobot(), data, meta)
}

func resourceRoboMakerRobotUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::Robot", ResourceRoboMakerRobot(), data, meta)
}

func resourceRoboMakerRobotDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::Robot", data, meta)
}

func resourceRoboMakerRobotCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}