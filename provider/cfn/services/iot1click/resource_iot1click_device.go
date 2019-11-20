// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioT1ClickDeviceType string = "AWS::IoT1Click::Device"

func ResourceIoT1ClickDevice() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoT1ClickDeviceRead,
		Create: resourceIoT1ClickDeviceCreate,
		Update: resourceIoT1ClickDeviceUpdate,
		Delete: resourceIoT1ClickDeviceDelete,
		CustomizeDiff: resourceIoT1ClickDeviceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoT1ClickDeviceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioT1ClickDeviceType, ResourceIoT1ClickDevice(), data, meta)
}

func resourceIoT1ClickDeviceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioT1ClickDeviceType, ResourceIoT1ClickDevice(), data, meta)
}

func resourceIoT1ClickDeviceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioT1ClickDeviceType, ResourceIoT1ClickDevice(), data, meta)
}

func resourceIoT1ClickDeviceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioT1ClickDeviceType, data, meta)
}

func resourceIoT1ClickDeviceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioT1ClickDeviceType, data, meta)
}
