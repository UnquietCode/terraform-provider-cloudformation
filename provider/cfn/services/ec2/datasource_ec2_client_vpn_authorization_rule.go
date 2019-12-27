// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnauthorizationrule.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnAuthorizationRuleType string = "AWS::EC2::ClientVpnAuthorizationRule"

func DatasourceEC2ClientVpnAuthorizationRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2ClientVpnAuthorizationRuleRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2ClientVpnAuthorizationRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnAuthorizationRuleType, DatasourceEC2ClientVpnAuthorizationRule(), data, meta)
}
