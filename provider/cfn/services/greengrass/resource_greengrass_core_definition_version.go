// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassCoreDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreengrassCoreDefinitionVersionCreate,
		Read:   resourceGreengrassCoreDefinitionVersionRead,
		Update: resourceGreengrassCoreDefinitionVersionUpdate,
		Delete: resourceGreengrassCoreDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"cores": {
				Type: schema.TypeList,
				Elem: propertyCore(),
				Required: true,
				ForceNew: true,
			},
			"core_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGreengrassCoreDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::CoreDefinitionVersion", data, meta)
}

func resourceGreengrassCoreDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::CoreDefinitionVersion", data, meta)
}

func resourceGreengrassCoreDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::CoreDefinitionVersion", data, meta)
}

func resourceGreengrassCoreDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::CoreDefinitionVersion", data, meta)
}