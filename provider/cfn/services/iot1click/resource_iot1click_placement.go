// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceIoT1ClickPlacementExists,
		Read:   resourceIoT1ClickPlacementRead,
		Create: resourceIoT1ClickPlacementCreate,
		Update: resourceIoT1ClickPlacementUpdate,
		Delete: resourceIoT1ClickPlacementDelete,
		CustomizeDiff: resourceIoT1ClickPlacementCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"placement_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"project_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"associated_devices": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"attributes": {
				Type: schema.TypeMap,
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

func resourceIoT1ClickPlacementExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoT1ClickPlacementRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Placement", ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Placement", ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Placement", ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Placement", data, meta)
}

func resourceIoT1ClickPlacementCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}