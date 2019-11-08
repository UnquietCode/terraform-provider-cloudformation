// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SpotFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SpotFleetCreate,
		Read:   resourceEC2SpotFleetRead,
		Update: resourceEC2SpotFleetUpdate,
		Delete: resourceEC2SpotFleetDelete,

		Schema: map[string]*schema.Schema{
			"spot_fleet_request_config_data": {
				Type: schema.TypeList,
				Elem: propertySpotFleetRequestConfigData(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceEC2SpotFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SpotFleet", data, meta)
}

func resourceEC2SpotFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SpotFleet", data, meta)
}

func resourceEC2SpotFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SpotFleet", data, meta)
}

func resourceEC2SpotFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SpotFleet", data, meta)
}