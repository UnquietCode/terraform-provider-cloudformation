// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SpotFleet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2SpotFleetExists,
		Read: resourceEC2SpotFleetRead,
		Create: resourceEC2SpotFleetCreate,
		Update: resourceEC2SpotFleetUpdate,
		Delete: resourceEC2SpotFleetDelete,
		CustomizeDiff: resourceEC2SpotFleetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"spot_fleet_request_config_data": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotFleetRequestConfigData(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SpotFleetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2SpotFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SpotFleet", ResourceEC2SpotFleet(), data, meta)
}

func resourceEC2SpotFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SpotFleet", ResourceEC2SpotFleet(), data, meta)
}

func resourceEC2SpotFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SpotFleet", ResourceEC2SpotFleet(), data, meta)
}

func resourceEC2SpotFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SpotFleet", data, meta)
}

func resourceEC2SpotFleetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
