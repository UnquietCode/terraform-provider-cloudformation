// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolidentityprovider.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolIdentityProviderExists,
		Read: resourceCognitoUserPoolIdentityProviderRead,
		Create: resourceCognitoUserPoolIdentityProviderCreate,
		Update: resourceCognitoUserPoolIdentityProviderUpdate,
		Delete: resourceCognitoUserPoolIdentityProviderDelete,
		CustomizeDiff: resourceCognitoUserPoolIdentityProviderCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"provider_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"attribute_mapping": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"provider_details": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"provider_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"idp_identifiers": {
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

func resourceCognitoUserPoolIdentityProviderExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolIdentityProviderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolIdentityProvider", ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolIdentityProvider", ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolIdentityProvider", ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolIdentityProvider", data, meta)
}

func resourceCognitoUserPoolIdentityProviderCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
