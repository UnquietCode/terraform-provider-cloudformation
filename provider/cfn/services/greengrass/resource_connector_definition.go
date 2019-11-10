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

func ResourceConnectorDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectorDefinitionCreate,
		Read:   resourceConnectorDefinitionRead,
		Update: resourceConnectorDefinitionUpdate,
		Delete: resourceConnectorDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyConnectorDefinitionConnectorDefinitionVersion(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConnectorDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceConnectorDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceConnectorDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceConnectorDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ConnectorDefinition", data, meta)
}