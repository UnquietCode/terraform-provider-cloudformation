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

func ResourceClientVpnTargetNetworkAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientVpnTargetNetworkAssociationCreate,
		Read:   resourceClientVpnTargetNetworkAssociationRead,
		Update: resourceClientVpnTargetNetworkAssociationUpdate,
		Delete: resourceClientVpnTargetNetworkAssociationDelete,

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

func resourceClientVpnTargetNetworkAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceClientVpnTargetNetworkAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceClientVpnTargetNetworkAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}

func resourceClientVpnTargetNetworkAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnTargetNetworkAssociation", data, meta)
}