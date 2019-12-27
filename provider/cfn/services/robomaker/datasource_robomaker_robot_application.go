// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceRoboMakerRobotApplication() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoboMakerRobotApplicationRead,
		
		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"robot_software_suite": {
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRoboMakerRobotApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(roboMakerRobotApplicationType, DatasourceRoboMakerRobotApplication(), data, meta)
}
