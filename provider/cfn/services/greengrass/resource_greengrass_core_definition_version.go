// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassCoreDefinitionVersionType string = "AWS::Greengrass::CoreDefinitionVersion"

func ResourceGreengrassCoreDefinitionVersion() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceGreengrassCoreDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassCoreDefinitionVersionType, ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassCoreDefinitionVersionType, ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassCoreDefinitionVersionType, ResourceGreengrassCoreDefinitionVersion(), data, meta)
}

func resourceGreengrassCoreDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassCoreDefinitionVersionType, data, meta)
}

func resourceGreengrassCoreDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassCoreDefinitionVersionType, data, meta)
}
