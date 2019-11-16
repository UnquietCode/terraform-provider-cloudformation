// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassConnectorDefinitionVersionType string = "AWS::Greengrass::ConnectorDefinitionVersion"

var greengrassConnectorDefinitionVersionProperties map[string]string = map[string]string{
	"connectors": "Connectors",
	"connector_definition_id": "ConnectorDefinitionId",
}

func ResourceGreengrassConnectorDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassConnectorDefinitionVersionExists,
		Read: resourceGreengrassConnectorDefinitionVersionRead,
		Create: resourceGreengrassConnectorDefinitionVersionCreate,
		Update: resourceGreengrassConnectorDefinitionVersionUpdate,
		Delete: resourceGreengrassConnectorDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassConnectorDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"connectors": {
				Type: schema.TypeList,
				Elem: propertyConnectorDefinitionVersionConnector(),
				Required: true,
			},
			"connector_definition_id": {
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

func resourceGreengrassConnectorDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassConnectorDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassConnectorDefinitionVersionType, ResourceGreengrassConnectorDefinitionVersion(), data, meta)
}

func resourceGreengrassConnectorDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassConnectorDefinitionVersionType, ResourceGreengrassConnectorDefinitionVersion(), data, greengrassConnectorDefinitionVersionProperties, meta)
}

func resourceGreengrassConnectorDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassConnectorDefinitionVersionType, ResourceGreengrassConnectorDefinitionVersion(), data, greengrassConnectorDefinitionVersionProperties, meta)
}

func resourceGreengrassConnectorDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassConnectorDefinitionVersionType, data, meta)
}

func resourceGreengrassConnectorDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassConnectorDefinitionVersionType, data, meta)
}
