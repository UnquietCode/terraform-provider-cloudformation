// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html

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
			"latest_version_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"the_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyCoreDefinitionCoreDefinitionVersion(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGreengrassCoreDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::CoreDefinition", ResourceGreengrassCoreDefinition(), data, meta)
}

func resourceGreengrassCoreDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::CoreDefinition", ResourceGreengrassCoreDefinition(), data, meta)
}

func resourceGreengrassCoreDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::CoreDefinition", ResourceGreengrassCoreDefinition(), data, meta)
}

func resourceGreengrassCoreDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::CoreDefinition", data, meta)
}