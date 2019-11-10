// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceInputCreate,
		Read:   resourceInputRead,
		Update: resourceInputUpdate,
		Delete: resourceInputDelete,

		Schema: map[string]*schema.Schema{
			"input_definition": {
				Type: schema.TypeList,
				Elem: propertyInputInputDefinition(),
				Required: false,
				MaxItems: 1,
			},
			"input_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"input_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTEvents::Input", data, meta)
}

func resourceInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTEvents::Input", data, meta)
}

func resourceInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTEvents::Input", data, meta)
}

func resourceInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTEvents::Input", data, meta)
}