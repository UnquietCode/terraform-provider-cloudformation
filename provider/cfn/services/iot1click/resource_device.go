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

func ResourceDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeviceCreate,
		Read:   resourceDeviceRead,
		Update: resourceDeviceUpdate,
		Delete: resourceDeviceDelete,

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

func resourceDeviceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Device", data, meta)
}

func resourceDeviceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Device", data, meta)
}

func resourceDeviceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Device", data, meta)
}

func resourceDeviceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Device", data, meta)
}