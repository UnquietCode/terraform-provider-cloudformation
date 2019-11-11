// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerRobotApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerRobotApplicationVersionExists,
		Read: resourceRoboMakerRobotApplicationVersionRead,
		Create: resourceRoboMakerRobotApplicationVersionCreate,
		Update: resourceRoboMakerRobotApplicationVersionUpdate,
		Delete: resourceRoboMakerRobotApplicationVersionDelete,
		CustomizeDiff: resourceRoboMakerRobotApplicationVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application": {
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

func resourceRoboMakerRobotApplicationVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerRobotApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplicationVersion", ResourceRoboMakerRobotApplicationVersion(), data, meta)
}

func resourceRoboMakerRobotApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplicationVersion", ResourceRoboMakerRobotApplicationVersion(), data, meta)
}

func resourceRoboMakerRobotApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplicationVersion", ResourceRoboMakerRobotApplicationVersion(), data, meta)
}

func resourceRoboMakerRobotApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRoboMakerRobotApplicationVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

