// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplication.html

package robomaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const roboMakerSimulationApplicationType string = "AWS::RoboMaker::SimulationApplication"

func ResourceRoboMakerSimulationApplication() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceRoboMakerSimulationApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, meta)
}

func resourceRoboMakerSimulationApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, meta)
}

func resourceRoboMakerSimulationApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(roboMakerSimulationApplicationType, ResourceRoboMakerSimulationApplication(), data, meta)
}

func resourceRoboMakerSimulationApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(roboMakerSimulationApplicationType, data, meta)
}

func resourceRoboMakerSimulationApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(roboMakerSimulationApplicationType, data, meta)
}
