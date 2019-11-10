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

func ResourceGreengrassLoggerDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassLoggerDefinitionVersionCreate,
		Read:   resourceGreengrassLoggerDefinitionVersionRead,
		Update: resourceGreengrassLoggerDefinitionVersionUpdate,
		Delete: resourceGreengrassLoggerDefinitionVersionDelete,

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

func resourceGreengrassLoggerDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceGreengrassLoggerDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceGreengrassLoggerDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}

func resourceGreengrassLoggerDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::LoggerDefinitionVersion", data, meta)
}