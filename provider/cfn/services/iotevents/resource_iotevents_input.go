// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTEventsInputType string = "AWS::IoTEvents::Input"

func ResourceIoTEventsInput() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoTEventsInputRead,
		Create: resourceIoTEventsInputCreate,
		Update: resourceIoTEventsInputUpdate,
		Delete: resourceIoTEventsInputDelete,
		CustomizeDiff: resourceIoTEventsInputCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"input_definition": {
				Type: schema.TypeSet,
				Elem: propertyInputInputDefinition(),
				Optional: true,
				MaxItems: 1,
			},
			"input_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTEventsInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTEventsInputType, ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTEventsInputType, ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTEventsInputType, ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTEventsInputType, data, meta)
}

func resourceIoTEventsInputCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTEventsInputType, data, meta)
}
