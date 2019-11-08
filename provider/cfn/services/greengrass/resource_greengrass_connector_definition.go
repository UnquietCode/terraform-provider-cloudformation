// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassConnectorDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassConnectorDefinitionCreate,
		Read:   resourceGreengrassConnectorDefinitionRead,
		Update: resourceGreengrassConnectorDefinitionUpdate,
		Delete: resourceGreengrassConnectorDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyConnectorDefinitionVersion(),
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

func resourceGreengrassConnectorDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceGreengrassConnectorDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceGreengrassConnectorDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ConnectorDefinition", data, meta)
}

func resourceGreengrassConnectorDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ConnectorDefinition", data, meta)
}