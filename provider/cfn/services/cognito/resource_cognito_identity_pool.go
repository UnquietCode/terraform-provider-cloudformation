// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoIdentityPoolType string = "AWS::Cognito::IdentityPool"

var cognitoIdentityPoolProperties map[string]string = map[string]string{
	"push_sync": "PushSync",
	"cognito_identity_providers": "CognitoIdentityProviders",
	"cognito_events": "CognitoEvents",
	"developer_provider_name": "DeveloperProviderName",
	"cognito_streams": "CognitoStreams",
	"identity_pool_name": "IdentityPoolName",
	"allow_unauthenticated_identities": "AllowUnauthenticatedIdentities",
	"supported_login_providers": "SupportedLoginProviders",
	"saml_provider_ar_ns": "SamlProviderARNs",
	"open_id_connect_provider_ar_ns": "OpenIdConnectProviderARNs",
}

func ResourceCognitoIdentityPool() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoIdentityPoolExists,
		Read: resourceCognitoIdentityPoolRead,
		Create: resourceCognitoIdentityPoolCreate,
		Update: resourceCognitoIdentityPoolUpdate,
		Delete: resourceCognitoIdentityPoolDelete,
		CustomizeDiff: resourceCognitoIdentityPoolCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"push_sync": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolPushSync(),
				Optional: true,
				MaxItems: 1,
			},
			"cognito_identity_providers": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolCognitoIdentityProvider(),
				Optional: true,
			},
			"cognito_events": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"developer_provider_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cognito_streams": {
				Type: schema.TypeList,
				Elem: propertyIdentityPoolCognitoStreams(),
				Optional: true,
				MaxItems: 1,
			},
			"identity_pool_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"allow_unauthenticated_identities": {
				Type: schema.TypeBool,
				Required: true,
			},
			"supported_login_providers": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"saml_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"open_id_connect_provider_ar_ns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceCognitoIdentityPoolExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoIdentityPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoIdentityPoolType, ResourceCognitoIdentityPool(), data, meta)
}

func resourceCognitoIdentityPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoIdentityPoolType, ResourceCognitoIdentityPool(), data, cognitoIdentityPoolProperties, meta)
}

func resourceCognitoIdentityPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoIdentityPoolType, ResourceCognitoIdentityPool(), data, cognitoIdentityPoolProperties, meta)
}

func resourceCognitoIdentityPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoIdentityPoolType, data, meta)
}

func resourceCognitoIdentityPoolCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoIdentityPoolType, data, meta)
}
