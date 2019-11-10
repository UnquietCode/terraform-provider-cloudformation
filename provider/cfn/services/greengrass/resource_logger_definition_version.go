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

func ResourceLoggerDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggerDefinitionVersionCreate,
		Read:   resourceLoggerDefinitionVersionRead,
		Update: resourceLoggerDefinitionVersionUpdate,
		Delete: resourceLoggerDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"logger_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loggers": {
				Type: schema.TypeList,
				Elem: propertyLoggerDefinitionVersionLogger(),
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLoggerDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceLoggerDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceLoggerDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceLoggerDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}