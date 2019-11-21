// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassLoggerDefinitionType string = "AWS::Greengrass::LoggerDefinition"

func ResourceGreengrassLoggerDefinition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassLoggerDefinitionRead,
		Create: resourceGreengrassLoggerDefinitionCreate,
		Update: resourceGreengrassLoggerDefinitionUpdate,
		Delete: resourceGreengrassLoggerDefinitionDelete,
		CustomizeDiff: resourceGreengrassLoggerDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertyLoggerDefinitionLoggerDefinitionVersion(),
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

func resourceGreengrassLoggerDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassLoggerDefinitionType, ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassLoggerDefinitionType, ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassLoggerDefinitionType, ResourceGreengrassLoggerDefinition(), data, meta)
}

func resourceGreengrassLoggerDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassLoggerDefinitionType, data, meta)
}

func resourceGreengrassLoggerDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassLoggerDefinitionType, data, meta)
}
