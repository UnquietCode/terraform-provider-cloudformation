// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoT1ClickDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoT1ClickDeviceCreate,
		Read:   resourceIoT1ClickDeviceRead,
		Update: resourceIoT1ClickDeviceUpdate,
		Delete: resourceIoT1ClickDeviceDelete,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceIoT1ClickDeviceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Device", data, meta)
}

func resourceIoT1ClickDeviceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Device", data, meta)
}

func resourceIoT1ClickDeviceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Device", data, meta)
}

func resourceIoT1ClickDeviceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Device", data, meta)
}