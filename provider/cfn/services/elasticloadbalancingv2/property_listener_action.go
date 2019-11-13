// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html

package elasticloadbalancingv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyListenerAction(extras...string) *schema.Resource {
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
			"authenticate_cognito_config": {
				Type: schema.TypeList,
				Elem: propertyListenerAuthenticateCognitoConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"authenticate_oidc_config": {
				Type: schema.TypeList,
				Elem: propertyListenerAuthenticateOidcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"fixed_response_config": {
				Type: schema.TypeList,
				Elem: propertyListenerFixedResponseConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"order": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"redirect_config": {
				Type: schema.TypeList,
				Elem: propertyListenerRedirectConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"target_group_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}