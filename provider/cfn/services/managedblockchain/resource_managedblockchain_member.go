// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const managedBlockchainMemberType string = "AWS::ManagedBlockchain::Member"

func ResourceManagedBlockchainMember() *schema.Resource {
	return &schema.Resource{
		Exists: resourceManagedBlockchainMemberExists,
		Read: resourceManagedBlockchainMemberRead,
		Create: resourceManagedBlockchainMemberCreate,
		Update: resourceManagedBlockchainMemberUpdate,
		Delete: resourceManagedBlockchainMemberDelete,
		CustomizeDiff: resourceManagedBlockchainMemberCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"member_configuration": {
				Type: schema.TypeSet,
				Elem: propertyMemberMemberConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"network_configuration": {
				Type: schema.TypeSet,
				Elem: propertyMemberNetworkConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"network_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"invitation_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceManagedBlockchainMemberExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceManagedBlockchainMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(managedBlockchainMemberType, ResourceManagedBlockchainMember(), data, meta)
}

func resourceManagedBlockchainMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(managedBlockchainMemberType, ResourceManagedBlockchainMember(), data, meta)
}

func resourceManagedBlockchainMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(managedBlockchainMemberType, ResourceManagedBlockchainMember(), data, meta)
}

func resourceManagedBlockchainMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(managedBlockchainMemberType, data, meta)
}

func resourceManagedBlockchainMemberCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(managedBlockchainMemberType, data, meta)
}
