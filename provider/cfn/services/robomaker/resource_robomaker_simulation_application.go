// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplication.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const roboMakerSimulationApplicationType string = "AWS::RoboMaker::SimulationApplication"

var roboMakerSimulationApplicationProperties map[string]string = map[string]string{
	"rendering_engine": "RenderingEngine",
	"simulation_software_suite": "SimulationSoftwareSuite",
	"current_revision_id": "CurrentRevisionId",
	"robot_software_suite": "RobotSoftwareSuite",
	"sources": "Sources",
	"tags": "Tags",
	"name": "Name",
}

func ResourceRoboMakerSimulationApplication() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerSimulationApplicationExists,
		Read: resourceRoboMakerSimulationApplicationRead,
		Create: resourceRoboMakerSimulationApplicationCreate,
		Update: resourceRoboMakerSimulationApplicationUpdate,
		Delete: resourceRoboMakerSimulationApplicationDelete,
		CustomizeDiff: resourceRoboMakerSimulationApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"rendering_engine": {
				Type: schema.TypeSet,
				Elem: propertySimulationApplicationRenderingEngine(),
				Required: true,
				MaxItems: 1,
			},
			"simulation_software_suite": {
				Type: schema.TypeSet,
				Elem: propertySimulationApplicationSimulationSoftwareSuite(),
				Required: true,
				MaxItems: 1,
			},
			"current_revision_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"robot_software_suite": {
				Type: schema.TypeSet,
				Elem: propertySimulationApplicationRobotSoftwareSuite(),
				Required: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertySimulationApplicationSourceConfig(),
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

func resourceRoboMakerSimulationApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerSimulationApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, meta)
}

func resourceRoboMakerSimulationApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, roboMakerSimulationApplicationProperties, meta)
}

func resourceRoboMakerSimulationApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, roboMakerSimulationApplicationProperties, meta)
}

func resourceRoboMakerSimulationApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(roboMakerSimulationApplicationType, data, meta)
}

func resourceRoboMakerSimulationApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(roboMakerSimulationApplicationType, data, meta)
}
