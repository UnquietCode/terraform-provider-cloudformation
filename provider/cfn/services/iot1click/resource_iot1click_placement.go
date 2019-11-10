// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoT1ClickPlacement() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoT1ClickPlacementCreate,
		Read:   resourceIoT1ClickPlacementRead,
		Update: resourceIoT1ClickPlacementUpdate,
		Delete: resourceIoT1ClickPlacementDelete,

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

func resourceIoT1ClickPlacementCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Placement", data, meta)
}

func resourceIoT1ClickPlacementRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Placement", data, meta)
}

func resourceIoT1ClickPlacementUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Placement", data, meta)
}

func resourceIoT1ClickPlacementDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Placement", data, meta)
}