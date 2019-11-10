// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceGreengrassResourceDefinitionCreate,
		Read:   resourceGreengrassResourceDefinitionRead,
		Update: resourceGreengrassResourceDefinitionUpdate,
		Delete: resourceGreengrassResourceDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"latest_version_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"id": {
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
				Computed: true,
			},
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionResourceDefinitionVersion(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceGreengrassResourceDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::ResourceDefinition", data, meta)
}

func resourceGreengrassResourceDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::ResourceDefinition", data, meta)
}

func resourceGreengrassResourceDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::ResourceDefinition", data, meta)
}

func resourceGreengrassResourceDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::ResourceDefinition", data, meta)
}