// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDeviceDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeviceDefinitionVersionCreate,
		Read:   resourceDeviceDefinitionVersionRead,
		Update: resourceDeviceDefinitionVersionUpdate,
		Delete: resourceDeviceDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"device_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"devices": {
				Type: schema.TypeList,
				Elem: propertyDeviceDefinitionVersionDevice(),
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDeviceDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::DeviceDefinitionVersion", data, meta)
}

func resourceDeviceDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::DeviceDefinitionVersion", data, meta)
}

func resourceDeviceDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::DeviceDefinitionVersion", data, meta)
}

func resourceDeviceDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::DeviceDefinitionVersion", data, meta)
}