// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassCoreDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassCoreDefinitionVersionExists,
		Read: resourceGreengrassCoreDefinitionVersionRead,
		Create: resourceGreengrassCoreDefinitionVersionCreate,
		Update: resourceGreengrassCoreDefinitionVersionUpdate,
		Delete: resourceGreengrassCoreDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassCoreDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cores": {
				Type: schema.TypeList,
				Elem: propertyCoreDefinitionVersionCore(),
				Required: true,
			},
			"core_definition_id": {
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

func resourceGreengrassCoreDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassCoreDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::CoreDefinitionVersion", ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::CoreDefinitionVersion", ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::CoreDefinitionVersion", ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::CoreDefinitionVersion", data, meta)
}

func resourceGreengrassCoreDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

