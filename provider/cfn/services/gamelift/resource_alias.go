// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliasCreate,
		Read:   resourceAliasRead,
		Update: resourceAliasUpdate,
		Delete: resourceAliasDelete,

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
				Elem: propertyAliasRoutingStrategy(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Alias", data, meta)
}

func resourceAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Alias", data, meta)
}

func resourceAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Alias", data, meta)
}

func resourceAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Alias", data, meta)
}