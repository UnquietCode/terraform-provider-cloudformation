// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2ClientVpnTargetNetworkAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2ClientVpnTargetNetworkAssociationCreate,
		Read:   resourceEC2ClientVpnTargetNetworkAssociationRead,
		Update: resourceEC2ClientVpnTargetNetworkAssociationUpdate,
		Delete: resourceEC2ClientVpnTargetNetworkAssociationDelete,

		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
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

func resourceEC2ClientVpnTargetNetworkAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}