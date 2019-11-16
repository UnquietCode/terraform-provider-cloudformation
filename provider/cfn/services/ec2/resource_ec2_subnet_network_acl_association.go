// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SubnetNetworkAclAssociationType string = "AWS::EC2::SubnetNetworkAclAssociation"

func ResourceEC2SubnetNetworkAclAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2SubnetNetworkAclAssociationExists,
		Read: resourceEC2SubnetNetworkAclAssociationRead,
		Create: resourceEC2SubnetNetworkAclAssociationCreate,
		Update: resourceEC2SubnetNetworkAclAssociationUpdate,
		Delete: resourceEC2SubnetNetworkAclAssociationDelete,
		CustomizeDiff: resourceEC2SubnetNetworkAclAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"network_acl_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SubnetNetworkAclAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2SubnetNetworkAclAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SubnetNetworkAclAssociationType, ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2SubnetNetworkAclAssociationType, ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2SubnetNetworkAclAssociationType, ResourceEC2SubnetNetworkAclAssociation(), data, meta)
}

func resourceEC2SubnetNetworkAclAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2SubnetNetworkAclAssociationType, data, meta)
}

func resourceEC2SubnetNetworkAclAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2SubnetNetworkAclAssociationType, data, meta)
}
