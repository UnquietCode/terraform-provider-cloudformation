// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html

package iot1click

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioT1ClickPlacementType string = "AWS::IoT1Click::Placement"

func ResourceIoT1ClickPlacement() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoT1ClickPlacementRead,
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"attributes": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func resourceIoT1ClickPlacementRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioT1ClickPlacementType, ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioT1ClickPlacementType, ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioT1ClickPlacementType, ResourceIoT1ClickPlacement(), data, meta)
}

func resourceIoT1ClickPlacementDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioT1ClickPlacementType, data, meta)
}

func resourceIoT1ClickPlacementCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioT1ClickPlacementType, data, meta)
}
