// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassFunctionDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassFunctionDefinitionExists,
		Read: resourceGreengrassFunctionDefinitionRead,
		Create: resourceGreengrassFunctionDefinitionCreate,
		Update: resourceGreengrassFunctionDefinitionUpdate,
		Delete: resourceGreengrassFunctionDefinitionDelete,
		CustomizeDiff: resourceGreengrassFunctionDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionFunctionDefinitionVersion(),
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

func resourceGreengrassFunctionDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassFunctionDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::FunctionDefinition", ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::FunctionDefinition", ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::FunctionDefinition", ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::FunctionDefinition", data, meta)
}

func resourceGreengrassFunctionDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
