// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolLambdaConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"create_auth_challenge": {
				Type: schema.TypeString,
				Required: false,
			},
			"pre_authentication": {
				Type: schema.TypeString,
				Required: false,
			},
			"define_auth_challenge": {
				Type: schema.TypeString,
				Required: false,
			},
			"pre_sign_up": {
				Type: schema.TypeString,
				Required: false,
			},
			"pre_token_generation": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_migration": {
				Type: schema.TypeString,
				Required: false,
			},
			"post_authentication": {
				Type: schema.TypeString,
				Required: false,
			},
			"post_confirmation": {
				Type: schema.TypeString,
				Required: false,
			},
			"custom_message": {
				Type: schema.TypeString,
				Required: false,
			},
			"verify_auth_challenge_response": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}