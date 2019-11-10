// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnauthorizationrule.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2ClientVpnAuthorizationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2ClientVpnAuthorizationRuleCreate,
		Read:   resourceEC2ClientVpnAuthorizationRuleRead,
		Delete: resourceEC2ClientVpnAuthorizationRuleDelete,

		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"access_group_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"target_network_cidr": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authorize_all_groups": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2ClientVpnAuthorizationRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnAuthorizationRule", ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnAuthorizationRule", ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnAuthorizationRule", ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnAuthorizationRule", data, meta)
}