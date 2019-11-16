// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassDeviceDefinitionVersionType string = "AWS::Greengrass::DeviceDefinitionVersion"

var greengrassDeviceDefinitionVersionProperties map[string]string = map[string]string{
	"device_definition_id": "DeviceDefinitionId",
	"devices": "Devices",
}

func ResourceGreengrassDeviceDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassDeviceDefinitionVersionExists,
		Read: resourceGreengrassDeviceDefinitionVersionRead,
		Create: resourceGreengrassDeviceDefinitionVersionCreate,
		Update: resourceGreengrassDeviceDefinitionVersionUpdate,
		Delete: resourceGreengrassDeviceDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassDeviceDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"device_definition_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"devices": {
				Type: schema.TypeList,
				Elem: propertyDeviceDefinitionVersionDevice(),
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

func resourceGreengrassDeviceDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassDeviceDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassDeviceDefinitionVersionType, ResourceGreengrassDeviceDefinitionVersion(), data, meta)
}

func resourceGreengrassDeviceDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassDeviceDefinitionVersionType, ResourceGreengrassDeviceDefinitionVersion(), data, greengrassDeviceDefinitionVersionProperties, meta)
}

func resourceGreengrassDeviceDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassDeviceDefinitionVersionType, ResourceGreengrassDeviceDefinitionVersion(), data, greengrassDeviceDefinitionVersionProperties, meta)
}

func resourceGreengrassDeviceDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassDeviceDefinitionVersionType, data, meta)
}

func resourceGreengrassDeviceDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassDeviceDefinitionVersionType, data, meta)
}
