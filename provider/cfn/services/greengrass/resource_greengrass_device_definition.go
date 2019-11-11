// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassDeviceDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassDeviceDefinitionExists,
		Read: resourceGreengrassDeviceDefinitionRead,
		Create: resourceGreengrassDeviceDefinitionCreate,
		Update: resourceGreengrassDeviceDefinitionUpdate,
		Delete: resourceGreengrassDeviceDefinitionDelete,
		CustomizeDiff: resourceGreengrassDeviceDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyDeviceDefinitionDeviceDefinitionVersion(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
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

func resourceGreengrassDeviceDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassDeviceDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::DeviceDefinition", ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::DeviceDefinition", ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::DeviceDefinition", ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::DeviceDefinition", data, meta)
}

func resourceGreengrassDeviceDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Greengrass::DeviceDefinition", data, meta)
}

