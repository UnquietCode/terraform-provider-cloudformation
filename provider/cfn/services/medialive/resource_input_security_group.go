// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInputSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceInputSecurityGroupCreate,
		Read:   resourceInputSecurityGroupRead,
		Update: resourceInputSecurityGroupUpdate,
		Delete: resourceInputSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"whitelist_rules": {
				Type: schema.TypeList,
				Elem: propertyInputSecurityGroupInputWhitelistRuleCidr(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceInputSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceInputSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceInputSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceInputSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::InputSecurityGroup", data, meta)
}