// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyListenerRuleAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authenticate_cognito_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleAuthenticateCognitoConfig(),
				Required: false,
				MaxItems: 1,
			},
			"authenticate_oidc_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleAuthenticateOidcConfig(),
				Required: false,
				MaxItems: 1,
			},
			"fixed_response_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleFixedResponseConfig(),
				Required: false,
				MaxItems: 1,
			},
			"order": {
				Type: schema.TypeInt,
				Required: false,
			},
			"redirect_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRuleRedirectConfig(),
				Required: false,
				MaxItems: 1,
			},
			"target_group_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}