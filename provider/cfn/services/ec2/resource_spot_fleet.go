// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSpotFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotFleetCreate,
		Read:   resourceSpotFleetRead,
		Update: resourceSpotFleetUpdate,
		Delete: resourceSpotFleetDelete,

		Schema: map[string]*schema.Schema{
			"spot_fleet_request_config_data": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotFleetRequestConfigData(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceSpotFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SpotFleet", data, meta)
}

func resourceSpotFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SpotFleet", data, meta)
}

func resourceSpotFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SpotFleet", data, meta)
}

func resourceSpotFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SpotFleet", data, meta)
}