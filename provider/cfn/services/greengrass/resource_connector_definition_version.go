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

func ResourceConnectorDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectorDefinitionVersionCreate,
		Read:   resourceConnectorDefinitionVersionRead,
		Update: resourceConnectorDefinitionVersionUpdate,
		Delete: resourceConnectorDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"connectors": {
				Type: schema.TypeList,
				Elem: propertyConnectorDefinitionVersionConnector(),
				Required: true,
				ForceNew: true,
			},
			"connector_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConnectorDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ConnectorDefinitionVersion", data, meta)
}

func resourceConnectorDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ConnectorDefinitionVersion", data, meta)
}

func resourceConnectorDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ConnectorDefinitionVersion", data, meta)
}

func resourceConnectorDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ConnectorDefinitionVersion", data, meta)
}