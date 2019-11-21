// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html

package greengrass

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const greengrassSubscriptionDefinitionType string = "AWS::Greengrass::SubscriptionDefinition"

func ResourceGreengrassSubscriptionDefinition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGreengrassSubscriptionDefinitionRead,
		Create: resourceGreengrassSubscriptionDefinitionCreate,
		Update: resourceGreengrassSubscriptionDefinitionUpdate,
		Delete: resourceGreengrassSubscriptionDefinitionDelete,
		CustomizeDiff: resourceGreengrassSubscriptionDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeSet,
				Elem: propertySubscriptionDefinitionSubscriptionDefinitionVersion(),
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

func resourceGreengrassSubscriptionDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(greengrassSubscriptionDefinitionType, ResourceGreengrassSubscriptionDefinition(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(greengrassSubscriptionDefinitionType, ResourceGreengrassSubscriptionDefinition(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(greengrassSubscriptionDefinitionType, ResourceGreengrassSubscriptionDefinition(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(greengrassSubscriptionDefinitionType, data, meta)
}

func resourceGreengrassSubscriptionDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(greengrassSubscriptionDefinitionType, data, meta)
}
