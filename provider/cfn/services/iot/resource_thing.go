// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceThingCreate,
		Read:   resourceThingRead,
		Update: resourceThingUpdate,
		Delete: resourceThingDelete,

		Schema: map[string]*schema.Schema{
			"attribute_payload": {
				Type: schema.TypeList,
				Elem: propertyThingAttributePayload(),
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

func resourceThingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::Thing", data, meta)
}

func resourceThingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::Thing", data, meta)
}

func resourceThingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::Thing", data, meta)
}

func resourceThingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::Thing", data, meta)
}