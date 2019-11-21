// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassFunctionDefinitionType string = "AWS::Greengrass::FunctionDefinition"

func ResourceGreengrassFunctionDefinition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassFunctionDefinitionRead,
		Create: resourceGreengrassFunctionDefinitionCreate,
		Update: resourceGreengrassFunctionDefinitionUpdate,
		Delete: resourceGreengrassFunctionDefinitionDelete,
		CustomizeDiff: resourceGreengrassFunctionDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertyFunctionDefinitionFunctionDefinitionVersion(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceGreengrassFunctionDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassFunctionDefinitionType, ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassFunctionDefinitionType, ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassFunctionDefinitionType, ResourceGreengrassFunctionDefinition(), data, meta)
}

func resourceGreengrassFunctionDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassFunctionDefinitionType, data, meta)
}

func resourceGreengrassFunctionDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassFunctionDefinitionType, data, meta)
}
