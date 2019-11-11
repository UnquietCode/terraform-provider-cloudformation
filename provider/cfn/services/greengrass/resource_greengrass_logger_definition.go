// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassLoggerDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassLoggerDefinitionExists,
		Read: resourceGreengrassLoggerDefinitionRead,
		Create: resourceGreengrassLoggerDefinitionCreate,
		Update: resourceGreengrassLoggerDefinitionUpdate,
		Delete: resourceGreengrassLoggerDefinitionDelete,
		CustomizeDiff: resourceGreengrassLoggerDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyLoggerDefinitionLoggerDefinitionVersion(),
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

func resourceGreengrassLoggerDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassLoggerDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::LoggerDefinition", ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::LoggerDefinition", ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::LoggerDefinition", ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::LoggerDefinition", data, meta)
}

func resourceGreengrassLoggerDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Greengrass::LoggerDefinition", data, meta)
}

