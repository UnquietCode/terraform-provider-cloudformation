// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnTargetNetworkAssociationType string = "AWS::EC2::ClientVpnTargetNetworkAssociation"

var eC2ClientVpnTargetNetworkAssociationProperties map[string]string = map[string]string{
	"client_vpn_endpoint_id": "ClientVpnEndpointId",
	"subnet_id": "SubnetId",
}

func ResourceEC2ClientVpnTargetNetworkAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2ClientVpnTargetNetworkAssociationExists,
		Read: resourceEC2ClientVpnTargetNetworkAssociationRead,
		Create: resourceEC2ClientVpnTargetNetworkAssociationCreate,
		Update: resourceEC2ClientVpnTargetNetworkAssociationUpdate,
		Delete: resourceEC2ClientVpnTargetNetworkAssociationDelete,
		CustomizeDiff: resourceEC2ClientVpnTargetNetworkAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
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

func resourceEC2ClientVpnTargetNetworkAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnTargetNetworkAssociationType, ResourceEC2ClientVpnTargetNetworkAssociation(), data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2ClientVpnTargetNetworkAssociationType, ResourceEC2ClientVpnTargetNetworkAssociation(), data, eC2ClientVpnTargetNetworkAssociationProperties, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2ClientVpnTargetNetworkAssociationType, ResourceEC2ClientVpnTargetNetworkAssociation(), data, eC2ClientVpnTargetNetworkAssociationProperties, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2ClientVpnTargetNetworkAssociationType, data, meta)
}

func resourceEC2ClientVpnTargetNetworkAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2ClientVpnTargetNetworkAssociationType, data, meta)
}
