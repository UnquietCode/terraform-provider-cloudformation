// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceIoTEventsInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTEventsInputCreate,
		Read:   resourceIoTEventsInputRead,
		Update: resourceIoTEventsInputUpdate,
		Delete: resourceIoTEventsInputDelete,

		Schema: map[string]*schema.Schema{
			"input_definition": {
				Type: schema.TypeList,
				Elem: propertyInputInputDefinition(),
				Optional: true,
				MaxItems: 1,
			},
			"input_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTEventsInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTEvents::Input", ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTEvents::Input", ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTEvents::Input", ResourceIoTEventsInput(), data, meta)
}

func resourceIoTEventsInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTEvents::Input", data, meta)
}