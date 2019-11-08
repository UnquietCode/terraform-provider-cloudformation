// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoIdentityPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceCognitoIdentityPoolCreate,
		Read:   resourceCognitoIdentityPoolRead,
		Update: resourceCognitoIdentityPoolUpdate,
		Delete: resourceCognitoIdentityPoolDelete,

		Schema: map[string]*schema.Schema{
			"push_sync": {
				Type: schema.TypeList,
				Elem: propertyPushSync(),
				Required: false,
				MaxItems: 1,
			},
			"cognito_identity_providers": {
				Type: schema.TypeList,
				Elem: propertyCognitoIdentityProvider(),
				Required: false,
			},
			"cognito_events": {
				Type: schema.TypeMap,
				Required: false,
			},
			"developer_provider_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"cognito_streams": {
				Type: schema.TypeList,
				Elem: propertyCognitoStreams(),
				Required: false,
				MaxItems: 1,
			},
			"identity_pool_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"allow_unauthenticated_identities": {
				Type: schema.TypeBool,
				Required: true,
			},
			"supported_login_providers": {
				Type: schema.TypeMap,
				Required: false,
			},
			"saml_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"open_id_connect_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceCognitoIdentityPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::IdentityPool", data, meta)
}

func resourceCognitoIdentityPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::IdentityPool", data, meta)
}

func resourceCognitoIdentityPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::IdentityPool", data, meta)
}

func resourceCognitoIdentityPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::IdentityPool", data, meta)
}