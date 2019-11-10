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

func ResourceSimulationApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimulationApplicationVersionCreate,
		Read:   resourceSimulationApplicationVersionRead,
		Update: resourceSimulationApplicationVersionUpdate,
		Delete: resourceSimulationApplicationVersionDelete,

		Schema: map[string]*schema.Schema{
			"current_revision_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"application": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSimulationApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RoboMaker::SimulationApplicationVersion", data, meta)
}

func resourceSimulationApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RoboMaker::SimulationApplicationVersion", data, meta)
}

func resourceSimulationApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RoboMaker::SimulationApplicationVersion", data, meta)
}

func resourceSimulationApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RoboMaker::SimulationApplicationVersion", data, meta)
}