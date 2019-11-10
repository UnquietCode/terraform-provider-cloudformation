// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPCDHCPOptionsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPCDHCPOptionsAssociationCreate,
		Read:   resourceEC2VPCDHCPOptionsAssociationRead,
		Update: resourceEC2VPCDHCPOptionsAssociationUpdate,
		Delete: resourceEC2VPCDHCPOptionsAssociationDelete,

		Schema: map[string]*schema.Schema{
			"dhcp_options_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPCDHCPOptionsAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCDHCPOptionsAssociation", ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCDHCPOptionsAssociation", ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCDHCPOptionsAssociation", ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCDHCPOptionsAssociation", data, meta)
}