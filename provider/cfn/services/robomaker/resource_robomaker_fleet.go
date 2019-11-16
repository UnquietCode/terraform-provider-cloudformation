// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-fleet.html

package robomaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const roboMakerFleetType string = "AWS::RoboMaker::Fleet"

func ResourceRoboMakerFleet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoboMakerFleetExists,
		Read: resourceRoboMakerFleetRead,
		Create: resourceRoboMakerFleetCreate,
		Update: resourceRoboMakerFleetUpdate,
		Delete: resourceRoboMakerFleetDelete,
		CustomizeDiff: resourceRoboMakerFleetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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

func resourceRoboMakerFleetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoboMakerFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(roboMakerFleetType, ResourceRoboMakerFleet(), data, meta)
}

func resourceRoboMakerFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(roboMakerFleetType, ResourceRoboMakerFleet(), data, meta)
}

func resourceRoboMakerFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(roboMakerFleetType, ResourceRoboMakerFleet(), data, meta)
}

func resourceRoboMakerFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(roboMakerFleetType, data, meta)
}

func resourceRoboMakerFleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(roboMakerFleetType, data, meta)
}
