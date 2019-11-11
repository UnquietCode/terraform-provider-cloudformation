// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerRobotApplication() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerRobotApplicationExists,
		Read: resourceRoboMakerRobotApplicationRead,
		Create: resourceRoboMakerRobotApplicationCreate,
		Update: resourceRoboMakerRobotApplicationUpdate,
		Delete: resourceRoboMakerRobotApplicationDelete,
		CustomizeDiff: resourceRoboMakerRobotApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"robot_software_suite": {
				Type: schema.TypeList,
				Elem: propertyRobotApplicationRobotSoftwareSuite(),
				Required: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyRobotApplicationSourceConfig(),
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

func resourceRoboMakerRobotApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerRobotApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplication", ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplication", ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplication", ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRoboMakerRobotApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::RoboMaker::RobotApplication", data, meta)
}

