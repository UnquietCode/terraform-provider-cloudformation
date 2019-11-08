// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerRobotApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoboMakerRobotApplicationVersionCreate,
		Read:   resourceRoboMakerRobotApplicationVersionRead,
		Update: resourceRoboMakerRobotApplicationVersionUpdate,
		Delete: resourceRoboMakerRobotApplicationVersionDelete,

		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"application": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRoboMakerRobotApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRoboMakerRobotApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRoboMakerRobotApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRoboMakerRobotApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}