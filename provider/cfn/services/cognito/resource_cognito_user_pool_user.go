// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolUser() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolUserExists,
		Read:   resourceCognitoUserPoolUserRead,
		Create: resourceCognitoUserPoolUserCreate,
		Update: resourceCognitoUserPoolUserUpdate,
		Delete: resourceCognitoUserPoolUserDelete,
		CustomizeDiff: resourceCognitoUserPoolUserCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"validation_data": {
				Type: schema.TypeList,
				Elem: propertyUserPoolUserAttributeType(),
				Optional: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"username": {
				Type: schema.TypeString,
				Optional: true,
			},
			"message_action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"desired_delivery_mediums": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"force_alias_creation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"user_attributes": {
				Type: schema.TypeList,
				Elem: propertyUserPoolUserAttributeType(),
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

func resourceCognitoUserPoolUserExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolUser", ResourceCognitoUserPoolUser(), data, meta)
}

func resourceCognitoUserPoolUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolUser", ResourceCognitoUserPoolUser(), data, meta)
}

func resourceCognitoUserPoolUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolUser", ResourceCognitoUserPoolUser(), data, meta)
}

func resourceCognitoUserPoolUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolUser", data, meta)
}

func resourceCognitoUserPoolUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}