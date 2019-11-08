// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMediaLiveInputSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediaLiveInputSecurityGroupCreate,
		Read:   resourceMediaLiveInputSecurityGroupRead,
		Update: resourceMediaLiveInputSecurityGroupUpdate,
		Delete: resourceMediaLiveInputSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"whitelist_rules": {
				Type: schema.TypeList,
				Elem: propertyInputWhitelistRuleCidr(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceMediaLiveInputSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceMediaLiveInputSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceMediaLiveInputSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::InputSecurityGroup", data, meta)
}

func resourceMediaLiveInputSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::InputSecurityGroup", data, meta)
}