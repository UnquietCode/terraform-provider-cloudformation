// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerRobot() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoboMakerRobotCreate,
		Read:   resourceRoboMakerRobotRead,
		Update: resourceRoboMakerRobotUpdate,
		Delete: resourceRoboMakerRobotDelete,

		Schema: map[string]*schema.Schema{
			"fleet": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"architecture": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"greengrass_group_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceRoboMakerRobotCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::Robot", data, meta)
}

func resourceRoboMakerRobotRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::Robot", data, meta)
}

func resourceRoboMakerRobotUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::Robot", data, meta)
}

func resourceRoboMakerRobotDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::Robot", data, meta)
}