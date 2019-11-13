// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassResourceDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassResourceDefinitionVersionExists,
		Read:   resourceGreengrassResourceDefinitionVersionRead,
		Create: resourceGreengrassResourceDefinitionVersionCreate,
		Update: resourceGreengrassResourceDefinitionVersionUpdate,
		Delete: resourceGreengrassResourceDefinitionVersionDelete,
		
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

func resourceGreengrassResourceDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassResourceDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ResourceDefinitionVersion", ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ResourceDefinitionVersion", ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ResourceDefinitionVersion", ResourceGreengrassResourceDefinitionVersion(), data, meta)
}

func resourceGreengrassResourceDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ResourceDefinitionVersion", data, meta)
}