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

func ResourceRoboMakerFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoboMakerFleetCreate,
		Read:   resourceRoboMakerFleetRead,
		Update: resourceRoboMakerFleetUpdate,
		Delete: resourceRoboMakerFleetDelete,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceRoboMakerFleetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::Fleet", data, meta)
}

func resourceRoboMakerFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::Fleet", data, meta)
}

func resourceRoboMakerFleetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::Fleet", data, meta)
}

func resourceRoboMakerFleetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::Fleet", data, meta)
}