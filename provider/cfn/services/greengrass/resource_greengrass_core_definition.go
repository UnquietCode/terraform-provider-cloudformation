// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassCoreDefinitionType string = "AWS::Greengrass::CoreDefinition"

var greengrassCoreDefinitionProperties map[string]string = map[string]string{
	"initial_version": "InitialVersion",
	"tags": "Tags",
	"name": "Name",
}

func ResourceGreengrassCoreDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassCoreDefinitionExists,
		Read: resourceGreengrassCoreDefinitionRead,
		Create: resourceGreengrassCoreDefinitionCreate,
		Update: resourceGreengrassCoreDefinitionUpdate,
		Delete: resourceGreengrassCoreDefinitionDelete,
		CustomizeDiff: resourceGreengrassCoreDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyCoreDefinitionCoreDefinitionVersion(),
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

func resourceGreengrassCoreDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassCoreDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassCoreDefinitionType, ResourceGreengrassCoreDefinition(), data, meta)
}

func resourceGreengrassCoreDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassCoreDefinitionType, ResourceGreengrassCoreDefinition(), data, greengrassCoreDefinitionProperties, meta)
}

func resourceGreengrassCoreDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassCoreDefinitionType, ResourceGreengrassCoreDefinition(), data, greengrassCoreDefinitionProperties, meta)
}

func resourceGreengrassCoreDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassCoreDefinitionType, data, meta)
}

func resourceGreengrassCoreDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassCoreDefinitionType, data, meta)
}
