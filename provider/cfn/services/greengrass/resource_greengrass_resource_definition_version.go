// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassResourceDefinitionVersionType string = "AWS::Greengrass::ResourceDefinitionVersion"

func ResourceGreengrassResourceDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassResourceDefinitionVersionRead,
		Create: resourceGreengrassResourceDefinitionVersionCreate,
		Update: resourceGreengrassResourceDefinitionVersionUpdate,
		Delete: resourceGreengrassResourceDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassResourceDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resources": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionResourceInstance(),
				Required: true,
			},
			"resource_definition_id": {
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

func resourceGreengrassResourceDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassResourceDefinitionVersionType, ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassResourceDefinitionVersionType, ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassResourceDefinitionVersionType, ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassResourceDefinitionVersionType, data, meta)
}

func resourceGreengrassResourceDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassResourceDefinitionVersionType, data, meta)
}
