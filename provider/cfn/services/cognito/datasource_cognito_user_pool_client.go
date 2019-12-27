// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolClientType string = "AWS::Cognito::UserPoolClient"

func DatasourceCognitoUserPoolClient() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCognitoUserPoolClientRead,
		
		Schema: map[string]*schema.Schema{
			"analytics_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolClientAnalyticsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"generate_secret": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"callback_ur_ls": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allowed_o_auth_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"read_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allowed_o_auth_flows_user_pool_client": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"default_redirect_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"supported_identity_providers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"client_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"allowed_o_auth_flows": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"explicit_auth_flows": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"logout_ur_ls": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"refresh_token_validity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"write_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"prevent_user_existence_errors": {
				Type: schema.TypeString,
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

func datasourceCognitoUserPoolClientRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolClientType, DatasourceCognitoUserPoolClient(), data, meta)
}
