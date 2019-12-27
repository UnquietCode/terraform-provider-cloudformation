// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html

package managedblockchain

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const managedBlockchainMemberType string = "AWS::ManagedBlockchain::Member"

func DatasourceManagedBlockchainMember() *schema.Resource {
	return &schema.Resource{
		Read: datasourceManagedBlockchainMemberRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceManagedBlockchainMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(managedBlockchainMemberType, DatasourceManagedBlockchainMember(), data, meta)
}
