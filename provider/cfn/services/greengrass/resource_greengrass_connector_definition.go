// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassConnectorDefinitionType string = "AWS::Greengrass::ConnectorDefinition"

func ResourceGreengrassConnectorDefinition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassConnectorDefinitionRead,
		Create: resourceGreengrassConnectorDefinitionCreate,
		Update: resourceGreengrassConnectorDefinitionUpdate,
		Delete: resourceGreengrassConnectorDefinitionDelete,
		CustomizeDiff: resourceGreengrassConnectorDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertyConnectorDefinitionConnectorDefinitionVersion(),
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

func resourceGreengrassConnectorDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassConnectorDefinitionType, ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassConnectorDefinitionType, ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassConnectorDefinitionType, ResourceGreengrassConnectorDefinition(), data, meta)
}

func resourceGreengrassConnectorDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassConnectorDefinitionType, data, meta)
}

func resourceGreengrassConnectorDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassConnectorDefinitionType, data, meta)
}
