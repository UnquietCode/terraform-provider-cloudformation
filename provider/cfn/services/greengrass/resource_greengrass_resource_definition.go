// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassResourceDefinitionType string = "AWS::Greengrass::ResourceDefinition"

func ResourceGreengrassResourceDefinition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassResourceDefinitionRead,
		Create: resourceGreengrassResourceDefinitionCreate,
		Update: resourceGreengrassResourceDefinitionUpdate,
		Delete: resourceGreengrassResourceDefinitionDelete,
		CustomizeDiff: resourceGreengrassResourceDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionResourceDefinitionVersion(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
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

func resourceGreengrassResourceDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassResourceDefinitionType, ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassResourceDefinitionType, ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassResourceDefinitionType, ResourceGreengrassResourceDefinition(), data, meta)
}

func resourceGreengrassResourceDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassResourceDefinitionType, data, meta)
}

func resourceGreengrassResourceDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassResourceDefinitionType, data, meta)
}
