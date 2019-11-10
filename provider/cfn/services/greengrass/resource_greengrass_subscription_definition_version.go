// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceGreengrassSubscriptionDefinitionVersionCreate,
		Read:   resourceGreengrassSubscriptionDefinitionVersionRead,
		Delete: resourceGreengrassSubscriptionDefinitionVersionDelete,

		Schema: map[string]*schema.Schema{
			"subscription_definition_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscriptions": {
				Type: schema.TypeList,
				Elem: propertySubscriptionDefinitionVersionSubscription(),
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGreengrassSubscriptionDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::SubscriptionDefinitionVersion", ResourceGreengrassSubscriptionDefinitionVersion(), data, meta)
}

func resourceGreengrassSubscriptionDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}