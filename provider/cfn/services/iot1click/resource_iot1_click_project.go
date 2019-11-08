// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot1click

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoT1ClickProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoT1ClickProjectCreate,
		Read:   resourceIoT1ClickProjectRead,
		Update: resourceIoT1ClickProjectUpdate,
		Delete: resourceIoT1ClickProjectDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"placement_template": {
				Type: schema.TypeList,
				Elem: propertyPlacementTemplate(),
				Required: true,
				MaxItems: 1,
			},
			"project_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceIoT1ClickProjectCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT1Click::Project", data, meta)
}

func resourceIoT1ClickProjectRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT1Click::Project", data, meta)
}

func resourceIoT1ClickProjectUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT1Click::Project", data, meta)
}

func resourceIoT1ClickProjectDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT1Click::Project", data, meta)
}