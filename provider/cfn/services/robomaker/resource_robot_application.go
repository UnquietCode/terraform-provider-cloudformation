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

func ResourceRobotApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceRobotApplicationCreate,
		Read:   resourceRobotApplicationRead,
		Update: resourceRobotApplicationUpdate,
		Delete: resourceRobotApplicationDelete,

		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"robot_software_suite": {
				Type: schema.TypeList,
				Elem: propertyRobotApplicationRobotSoftwareSuite(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyRobotApplicationSourceConfig(),
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

func resourceRobotApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRobotApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRobotApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplication", data, meta)
}

func resourceRobotApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplication", data, meta)
}