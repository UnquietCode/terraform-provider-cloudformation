// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SubnetNetworkAclAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SubnetNetworkAclAssociationCreate,
		Read:   resourceEC2SubnetNetworkAclAssociationRead,
		Delete: resourceEC2SubnetNetworkAclAssociationDelete,

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"network_acl_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SubnetNetworkAclAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetNetworkAclAssociation", ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetNetworkAclAssociation", ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetNetworkAclAssociation", ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetNetworkAclAssociation", data, meta)
}