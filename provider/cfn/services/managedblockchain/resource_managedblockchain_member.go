// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceManagedBlockchainMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceManagedBlockchainMemberCreate,
		Read:   resourceManagedBlockchainMemberRead,
		Update: resourceManagedBlockchainMemberUpdate,
		Delete: resourceManagedBlockchainMemberDelete,

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

func resourceManagedBlockchainMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceManagedBlockchainMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceManagedBlockchainMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ManagedBlockchain::Member", data, meta)
}

func resourceManagedBlockchainMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ManagedBlockchain::Member", data, meta)
}