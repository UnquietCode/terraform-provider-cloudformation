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

func ResourceRoboMakerRobotApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoboMakerRobotApplicationCreate,
		Read:   resourceRoboMakerRobotApplicationRead,
		Update: resourceRoboMakerRobotApplicationUpdate,
		Delete: resourceRoboMakerRobotApplicationDelete,

		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"robot_software_suite": {
				Type: schema.TypeList,
				Elem: propertyRobotSoftwareSuite(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertySourceConfig(),
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceRoboMakerRobotApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRoboMakerRobotApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRoboMakerRobotApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRoboMakerRobotApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplication", data, meta)
}