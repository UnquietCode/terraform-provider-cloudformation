// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html

package elasticloadbalancingv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var listenerAuthenticateOidcConfigProperties map[string]string = map[string]string{
	"authentication_request_extra_params": "AuthenticationRequestExtraParams",
	"authorization_endpoint": "AuthorizationEndpoint",
	"client_id": "ClientId",
	"client_secret": "ClientSecret",
	"issuer": "Issuer",
	"on_unauthenticated_request": "OnUnauthenticatedRequest",
	"scope": "Scope",
	"session_cookie_name": "SessionCookieName",
	"session_timeout": "SessionTimeout",
	"token_endpoint": "TokenEndpoint",
	"user_info_endpoint": "UserInfoEndpoint",
}

func propertyListenerAuthenticateOidcConfig(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authentication_request_extra_params": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"authorization_endpoint": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type: schema.TypeString,
				Required: true,
			},
			"issuer": {
				Type: schema.TypeString,
				Required: true,
			},
			"on_unauthenticated_request": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type: schema.TypeString,
				Optional: true,
			},
			"session_cookie_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"session_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"token_endpoint": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_info_endpoint": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
