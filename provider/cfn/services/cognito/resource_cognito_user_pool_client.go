// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceCognitoUserPoolClientCreate,
		Read:   resourceCognitoUserPoolClientRead,
		Update: resourceCognitoUserPoolClientUpdate,
		Delete: resourceCognitoUserPoolClientDelete,

		Schema: map[string]*schema.Schema{
			"analytics_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolClientAnalyticsConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"generate_secret": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"callback_ur_ls": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"allowed_o_auth_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"read_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"allowed_o_auth_flows_user_pool_client": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"default_redirect_uri": {
				Type: schema.TypeString,
				Required: false,
			},
			"supported_identity_providers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"client_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowed_o_auth_flows": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"explicit_auth_flows": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"logout_ur_ls": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"refresh_token_validity": {
				Type: schema.TypeInt,
				Required: false,
			},
			"write_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceCognitoUserPoolClientCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolClient", data, meta)
}

func resourceCognitoUserPoolClientRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolClient", data, meta)
}

func resourceCognitoUserPoolClientUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolClient", data, meta)
}

func resourceCognitoUserPoolClientDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolClient", data, meta)
}