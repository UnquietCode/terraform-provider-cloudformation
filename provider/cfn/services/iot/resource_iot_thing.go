// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTThingType string = "AWS::IoT::Thing"

var ioTThingProperties map[string]string = map[string]string{
	"attribute_payload": "AttributePayload",
	"thing_name": "ThingName",
}

func ResourceIoTThing() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTThingExists,
		Read: resourceIoTThingRead,
		Create: resourceIoTThingCreate,
		Update: resourceIoTThingUpdate,
		Delete: resourceIoTThingDelete,
		CustomizeDiff: resourceIoTThingCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"attribute_payload": {
				Type: schema.TypeSet,
				Elem: propertyThingAttributePayload(),
				Optional: true,
				MaxItems: 1,
			},
			"thing_name": {
				Type: schema.TypeString,
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

func resourceIoTThingExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTThingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTThingType, ResourceIoTThing(), data, meta)
}

func resourceIoTThingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTThingType, ResourceIoTThing(), data, ioTThingProperties, meta)
}

func resourceIoTThingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTThingType, ResourceIoTThing(), data, ioTThingProperties, meta)
}

func resourceIoTThingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTThingType, data, meta)
}

func resourceIoTThingCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTThingType, data, meta)
}
