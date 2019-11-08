// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTThingCreate,
		Read:   resourceIoTThingRead,
		Update: resourceIoTThingUpdate,
		Delete: resourceIoTThingDelete,

		Schema: map[string]*schema.Schema{
			"attribute_payload": {
				Type: schema.TypeList,
				Elem: propertyAttributePayload(),
				Required: false,
				MaxItems: 1,
			},
			"thing_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTThingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::Thing", data, meta)
}

func resourceIoTThingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::Thing", data, meta)
}

func resourceIoTThingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::Thing", data, meta)
}

func resourceIoTThingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::Thing", data, meta)
}