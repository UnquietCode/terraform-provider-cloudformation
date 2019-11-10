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

func ResourceClientVpnAuthorizationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientVpnAuthorizationRuleCreate,
		Read:   resourceClientVpnAuthorizationRuleRead,
		Update: resourceClientVpnAuthorizationRuleUpdate,
		Delete: resourceClientVpnAuthorizationRuleDelete,

		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"access_group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"target_network_cidr": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authorize_all_groups": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceClientVpnAuthorizationRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnAuthorizationRule", data, meta)
}

func resourceClientVpnAuthorizationRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnAuthorizationRule", data, meta)
}

func resourceClientVpnAuthorizationRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnAuthorizationRule", data, meta)
}

func resourceClientVpnAuthorizationRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnAuthorizationRule", data, meta)
}