// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassFunctionDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassFunctionDefinitionVersionExists,
		Read: resourceGreengrassFunctionDefinitionVersionRead,
		Create: resourceGreengrassFunctionDefinitionVersionCreate,
		Update: resourceGreengrassFunctionDefinitionVersionUpdate,
		Delete: resourceGreengrassFunctionDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassFunctionDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"default_config": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionVersionDefaultConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"functions": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionVersionFunction(),
				Required: true,
			},
			"function_definition_id": {
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

func resourceGreengrassFunctionDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassFunctionDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::FunctionDefinitionVersion", ResourceGreengrassFunctionDefinitionVersion(), data, meta)
}

func resourceGreengrassFunctionDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::FunctionDefinitionVersion", ResourceGreengrassFunctionDefinitionVersion(), data, meta)
}

func resourceGreengrassFunctionDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::FunctionDefinitionVersion", ResourceGreengrassFunctionDefinitionVersion(), data, meta)
}

func resourceGreengrassFunctionDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::FunctionDefinitionVersion", data, meta)
}

func resourceGreengrassFunctionDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Greengrass::FunctionDefinitionVersion", data, meta)
}

