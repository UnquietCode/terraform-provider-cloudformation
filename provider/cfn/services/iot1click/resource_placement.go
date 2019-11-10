// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePlacement() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlacementCreate,
		Read:   resourcePlacementRead,
		Update: resourcePlacementUpdate,
		Delete: resourcePlacementDelete,

		Schema: map[string]*schema.Schema{
			"placement_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"project_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"associated_devices": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"attributes": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourcePlacementCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Placement", data, meta)
}

func resourcePlacementRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Placement", data, meta)
}

func resourcePlacementUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Placement", data, meta)
}

func resourcePlacementDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Placement", data, meta)
}