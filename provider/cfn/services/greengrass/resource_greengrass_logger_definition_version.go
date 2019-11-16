// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassLoggerDefinitionVersionType string = "AWS::Greengrass::LoggerDefinitionVersion"

func ResourceGreengrassLoggerDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassLoggerDefinitionVersionExists,
		Read: resourceGreengrassLoggerDefinitionVersionRead,
		Create: resourceGreengrassLoggerDefinitionVersionCreate,
		Update: resourceGreengrassLoggerDefinitionVersionUpdate,
		Delete: resourceGreengrassLoggerDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassLoggerDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"logger_definition_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"loggers": {
				Type: schema.TypeList,
				Elem: propertyLoggerDefinitionVersionLogger(),
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

func resourceGreengrassLoggerDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassLoggerDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassLoggerDefinitionVersionType, ResourceGreengrassLoggerDefinitionVersion(), data, meta)
}

func resourceGreengrassLoggerDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassLoggerDefinitionVersionType, ResourceGreengrassLoggerDefinitionVersion(), data, meta)
}

func resourceGreengrassLoggerDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassLoggerDefinitionVersionType, ResourceGreengrassLoggerDefinitionVersion(), data, meta)
}

func resourceGreengrassLoggerDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassLoggerDefinitionVersionType, data, meta)
}

func resourceGreengrassLoggerDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassLoggerDefinitionVersionType, data, meta)
}
