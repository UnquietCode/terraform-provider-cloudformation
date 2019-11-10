// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemberCreate,
		Read:   resourceMemberRead,
		Update: resourceMemberUpdate,
		Delete: resourceMemberDelete,

		Schema: map[string]*schema.Schema{
			"member_configuration": {
				Type: schema.TypeList,
				Elem: propertyMemberMemberConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyMemberNetworkConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"network_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"invitation_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ManagedBlockchain::Member", data, meta)
}