// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html

package robomaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const roboMakerRobotApplicationType string = "AWS::RoboMaker::RobotApplication"

func ResourceRoboMakerRobotApplication() *schema.Resource {
	return &schema.Resource{
		Read: resourceRoboMakerRobotApplicationRead,
		Create: resourceRoboMakerRobotApplicationCreate,
		Update: resourceRoboMakerRobotApplicationUpdate,
		Delete: resourceRoboMakerRobotApplicationDelete,
		CustomizeDiff: resourceRoboMakerRobotApplicationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"robot_software_suite": {
				Type: schema.TypeSet,
				Elem: propertyRobotApplicationRobotSoftwareSuite(),
				Required: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyRobotApplicationSourceConfig(),
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

func resourceRoboMakerRobotApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(roboMakerRobotApplicationType, ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(roboMakerRobotApplicationType, ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(roboMakerRobotApplicationType, ResourceRoboMakerRobotApplication(), data, meta)
}

func resourceRoboMakerRobotApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(roboMakerRobotApplicationType, data, meta)
}

func resourceRoboMakerRobotApplicationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(roboMakerRobotApplicationType, data, meta)
}
