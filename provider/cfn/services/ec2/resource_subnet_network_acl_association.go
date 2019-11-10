// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSubnetNetworkAclAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetNetworkAclAssociationCreate,
		Read:   resourceSubnetNetworkAclAssociationRead,
		Update: resourceSubnetNetworkAclAssociationUpdate,
		Delete: resourceSubnetNetworkAclAssociationDelete,

		Schema: map[string]*schema.Schema{
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

func resourceSubnetNetworkAclAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetNetworkAclAssociation", data, meta)
}

func resourceSubnetNetworkAclAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetNetworkAclAssociation", data, meta)
}

func resourceSubnetNetworkAclAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetNetworkAclAssociation", data, meta)
}

func resourceSubnetNetworkAclAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetNetworkAclAssociation", data, meta)
}