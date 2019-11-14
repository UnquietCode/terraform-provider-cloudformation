// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGreengrassSubscriptionDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGreengrassSubscriptionDefinitionVersionExists,
		Read: resourceGreengrassSubscriptionDefinitionVersionRead,
		Create: resourceGreengrassSubscriptionDefinitionVersionCreate,
		Update: resourceGreengrassSubscriptionDefinitionVersionUpdate,
		Delete: resourceGreengrassSubscriptionDefinitionVersionDelete,
		CustomizeDiff: resourceGreengrassSubscriptionDefinitionVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"subscription_definition_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subscriptions": {
				Type: schema.TypeList,
				Elem: propertySubscriptionDefinitionVersionSubscription(),
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

func resourceGreengrassSubscriptionDefinitionVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

