// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolidentityprovider.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolIdentityProviderType string = "AWS::Cognito::UserPoolIdentityProvider"

func ResourceCognitoUserPoolIdentityProvider() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"provider_details": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCognitoUserPoolIdentityProviderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolIdentityProviderType, ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolIdentityProviderType, ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolIdentityProviderType, ResourceCognitoUserPoolIdentityProvider(), data, meta)
}

func resourceCognitoUserPoolIdentityProviderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolIdentityProviderType, data, meta)
}

func resourceCognitoUserPoolIdentityProviderCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolIdentityProviderType, data, meta)
}
