// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGameLiftAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameLiftAliasCreate,
		Read:   resourceGameLiftAliasRead,
		Update: resourceGameLiftAliasUpdate,
		Delete: resourceGameLiftAliasDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"routing_strategy": {
				Type: schema.TypeList,
				Elem: propertyRoutingStrategy(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceGameLiftAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Alias", data, meta)
}

func resourceGameLiftAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Alias", data, meta)
}

func resourceGameLiftAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Alias", data, meta)
}

func resourceGameLiftAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Alias", data, meta)
}