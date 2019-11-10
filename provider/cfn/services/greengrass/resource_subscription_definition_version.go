// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSubscriptionDefinitionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubscriptionDefinitionVersionCreate,
		Read:   resourceSubscriptionDefinitionVersionRead,
		Update: resourceSubscriptionDefinitionVersionUpdate,
		Delete: resourceSubscriptionDefinitionVersionDelete,

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

func resourceSubscriptionDefinitionVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}

func resourceSubscriptionDefinitionVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}

func resourceSubscriptionDefinitionVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}

func resourceSubscriptionDefinitionVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::SubscriptionDefinitionVersion", data, meta)
}