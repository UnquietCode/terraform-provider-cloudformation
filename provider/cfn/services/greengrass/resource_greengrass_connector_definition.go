// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassConnectorDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassConnectorDefinitionExists,
		Read: resourceGreengrassConnectorDefinitionRead,
		Create: resourceGreengrassConnectorDefinitionCreate,
		Update: resourceGreengrassConnectorDefinitionUpdate,
		Delete: resourceGreengrassConnectorDefinitionDelete,
		CustomizeDiff: resourceGreengrassConnectorDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyConnectorDefinitionConnectorDefinitionVersion(),
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

func resourceGreengrassConnectorDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassConnectorDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ConnectorDefinition", ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ConnectorDefinition", ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ConnectorDefinition", ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceGreengrassConnectorDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Greengrass::ConnectorDefinition", data, meta)
}

