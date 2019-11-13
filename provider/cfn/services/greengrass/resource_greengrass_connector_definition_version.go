// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassConnectorDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassConnectorDefinitionVersionExists,
		Read:   resourceGreengrassConnectorDefinitionVersionRead,
		Create: resourceGreengrassConnectorDefinitionVersionCreate,
		Update: resourceGreengrassConnectorDefinitionVersionUpdate,
		Delete: resourceGreengrassConnectorDefinitionVersionDelete,
		
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
	return plugin.ResourceRead("AWS::Greengrass::ConnectorDefinitionVersion", ResourceGreengrassConnectorDefinitionVersion(), data, meta)
}

func resourceGreengrassConnectorDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ConnectorDefinitionVersion", ResourceGreengrassConnectorDefinitionVersion(), data, meta)
}

func resourceGreengrassConnectorDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ConnectorDefinitionVersion", ResourceGreengrassConnectorDefinitionVersion(), data, meta)
}

func resourceGreengrassConnectorDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ConnectorDefinitionVersion", data, meta)
}