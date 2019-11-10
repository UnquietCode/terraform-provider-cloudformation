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

func ResourceLoggerDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggerDefinitionCreate,
		Read:   resourceLoggerDefinitionRead,
		Update: resourceLoggerDefinitionUpdate,
		Delete: resourceLoggerDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyLoggerDefinitionLoggerDefinitionVersion(),
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

func resourceLoggerDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::LoggerDefinition", data, meta)
}

func resourceLoggerDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::LoggerDefinition", data, meta)
}

func resourceLoggerDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::LoggerDefinition", data, meta)
}

func resourceLoggerDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::LoggerDefinition", data, meta)
}