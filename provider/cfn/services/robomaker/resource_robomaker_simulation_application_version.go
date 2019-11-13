// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoboMakerSimulationApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerSimulationApplicationVersionExists,
		Read:   resourceRoboMakerSimulationApplicationVersionRead,
		Create: resourceRoboMakerSimulationApplicationVersionCreate,
		Update: resourceRoboMakerSimulationApplicationVersionUpdate,
		Delete: resourceRoboMakerSimulationApplicationVersionDelete,
		CustomizeDiff: resourceRoboMakerSimulationApplicationVersionCustomizeDiff,

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

func resourceRoboMakerSimulationApplicationVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerSimulationApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::SimulationApplicationVersion", ResourceRoboMakerSimulationApplicationVersion(), data, meta)
}

func resourceRoboMakerSimulationApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::SimulationApplicationVersion", ResourceRoboMakerSimulationApplicationVersion(), data, meta)
}

func resourceRoboMakerSimulationApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::SimulationApplicationVersion", ResourceRoboMakerSimulationApplicationVersion(), data, meta)
}

func resourceRoboMakerSimulationApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::SimulationApplicationVersion", data, meta)
}

func resourceRoboMakerSimulationApplicationVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}