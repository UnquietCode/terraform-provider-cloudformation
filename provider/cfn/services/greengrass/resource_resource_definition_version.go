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

func ResourceResourceDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourceDefinitionVersionCreate,
		Read:   resourceResourceDefinitionVersionRead,
		Update: resourceResourceDefinitionVersionUpdate,
		Delete: resourceResourceDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"resources": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionResourceInstance(),
				Required: true,
				ForceNew: true,
			},
			"resource_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceResourceDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ResourceDefinitionVersion", data, meta)
}

func resourceResourceDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ResourceDefinitionVersion", data, meta)
}

func resourceResourceDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ResourceDefinitionVersion", data, meta)
}

func resourceResourceDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ResourceDefinitionVersion", data, meta)
}