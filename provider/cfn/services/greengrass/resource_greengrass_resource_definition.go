// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassResourceDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassResourceDefinitionExists,
		Read: resourceGreengrassResourceDefinitionRead,
		Create: resourceGreengrassResourceDefinitionCreate,
		Update: resourceGreengrassResourceDefinitionUpdate,
		Delete: resourceGreengrassResourceDefinitionDelete,
		CustomizeDiff: resourceGreengrassResourceDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionResourceDefinitionVersion(),
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

func resourceGreengrassResourceDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassResourceDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ResourceDefinition", ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ResourceDefinition", ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ResourceDefinition", ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ResourceDefinition", data, meta)
}

func resourceGreengrassResourceDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
