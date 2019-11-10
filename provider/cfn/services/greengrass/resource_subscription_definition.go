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

func ResourceSubscriptionDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubscriptionDefinitionCreate,
		Read:   resourceSubscriptionDefinitionRead,
		Update: resourceSubscriptionDefinitionUpdate,
		Delete: resourceSubscriptionDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"initial_version": {
				Type: schema.TypeList,
				Elem: propertySubscriptionDefinitionSubscriptionDefinitionVersion(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSubscriptionDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Greengrass::SubscriptionDefinition", data, meta)
}

func resourceSubscriptionDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Greengrass::SubscriptionDefinition", data, meta)
}

func resourceSubscriptionDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Greengrass::SubscriptionDefinition", data, meta)
}

func resourceSubscriptionDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Greengrass::SubscriptionDefinition", data, meta)
}