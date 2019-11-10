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

func ResourceRobotApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceRobotApplicationVersionCreate,
		Read:   resourceRobotApplicationVersionRead,
		Update: resourceRobotApplicationVersionUpdate,
		Delete: resourceRobotApplicationVersionDelete,

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

func resourceRobotApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRobotApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRobotApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}

func resourceRobotApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::RobotApplicationVersion", data, meta)
}