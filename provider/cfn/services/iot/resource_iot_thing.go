// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html

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
				Elem: propertyThingAttributePayload(),
				Optional: true,
				MaxItems: 1,
			},
			"thing_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTThingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::Thing", ResourceIoTThing(), data, meta)
}

func resourceIoTThingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::Thing", ResourceIoTThing(), data, meta)
}

func resourceIoTThingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::Thing", ResourceIoTThing(), data, meta)
}

func resourceIoTThingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::Thing", data, meta)
}