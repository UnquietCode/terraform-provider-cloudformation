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

func ResourceGreengrassCoreDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassCoreDefinitionCreate,
		Read:   resourceGreengrassCoreDefinitionRead,
		Update: resourceGreengrassCoreDefinitionUpdate,
		Delete: resourceGreengrassCoreDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyCoreDefinitionCoreDefinitionVersion(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGreengrassCoreDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::CoreDefinition", data, meta)
}

func resourceGreengrassCoreDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::CoreDefinition", data, meta)
}

func resourceGreengrassCoreDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::CoreDefinition", data, meta)
}

func resourceGreengrassCoreDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::CoreDefinition", data, meta)
}