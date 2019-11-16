// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnauthorizationrule.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnAuthorizationRuleType string = "AWS::EC2::ClientVpnAuthorizationRule"

func ResourceEC2ClientVpnAuthorizationRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2ClientVpnAuthorizationRuleExists,
		Read: resourceEC2ClientVpnAuthorizationRuleRead,
		Create: resourceEC2ClientVpnAuthorizationRuleCreate,
		Update: resourceEC2ClientVpnAuthorizationRuleUpdate,
		Delete: resourceEC2ClientVpnAuthorizationRuleDelete,
		CustomizeDiff: resourceEC2ClientVpnAuthorizationRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"access_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_network_cidr": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorize_all_groups": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2ClientVpnAuthorizationRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnAuthorizationRuleType, ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2ClientVpnAuthorizationRuleType, ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2ClientVpnAuthorizationRuleType, ResourceEC2ClientVpnAuthorizationRule(), data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2ClientVpnAuthorizationRuleType, data, meta)
}

func resourceEC2ClientVpnAuthorizationRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2ClientVpnAuthorizationRuleType, data, meta)
}
