// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassDeviceDefinitionType string = "AWS::Greengrass::DeviceDefinition"

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
				Type: schema.TypeSet,
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
	return plugin.ResourceRead(greengrassDeviceDefinitionType, ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassDeviceDefinitionType, ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassDeviceDefinitionType, ResourceGreengrassDeviceDefinition(), data, meta)
}

func resourceGreengrassDeviceDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassDeviceDefinitionType, data, meta)
}

func resourceGreengrassDeviceDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassDeviceDefinitionType, data, meta)
}
