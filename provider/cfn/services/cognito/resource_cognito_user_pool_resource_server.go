// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolResourceServer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolResourceServerExists,
		Read: resourceCognitoUserPoolResourceServerRead,
		Create: resourceCognitoUserPoolResourceServerCreate,
		Update: resourceCognitoUserPoolResourceServerUpdate,
		Delete: resourceCognitoUserPoolResourceServerDelete,
		CustomizeDiff: resourceCognitoUserPoolResourceServerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"scopes": {
				Type: schema.TypeList,
				Elem: propertyUserPoolResourceServerResourceServerScopeType(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCognitoUserPoolResourceServerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolResourceServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolResourceServer", ResourceCognitoUserPoolResourceServer(), data, meta)
}

func resourceCognitoUserPoolResourceServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolResourceServer", ResourceCognitoUserPoolResourceServer(), data, meta)
}

func resourceCognitoUserPoolResourceServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolResourceServer", ResourceCognitoUserPoolResourceServer(), data, meta)
}

func resourceCognitoUserPoolResourceServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolResourceServer", data, meta)
}

func resourceCognitoUserPoolResourceServerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
