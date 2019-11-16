// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolLambdaConfig(extras...string) *schema.Resource {
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
			"create_auth_challenge": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pre_authentication": {
				Type: schema.TypeString,
				Optional: true,
			},
			"define_auth_challenge": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pre_sign_up": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pre_token_generation": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_migration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"post_authentication": {
				Type: schema.TypeString,
				Optional: true,
			},
			"post_confirmation": {
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"verify_auth_challenge_response": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
