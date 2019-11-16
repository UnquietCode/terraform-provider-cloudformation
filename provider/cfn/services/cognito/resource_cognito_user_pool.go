// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolType string = "AWS::Cognito::UserPool"

var cognitoUserPoolProperties map[string]string = map[string]string{
	"user_pool_tags": "UserPoolTags",
	"policies": "Policies",
	"verification_message_template": "VerificationMessageTemplate",
	"mfa_configuration": "MfaConfiguration",
	"schema": "Schema",
	"admin_create_user_config": "AdminCreateUserConfig",
	"sms_authentication_message": "SmsAuthenticationMessage",
	"user_pool_name": "UserPoolName",
	"sms_verification_message": "SmsVerificationMessage",
	"user_pool_add_ons": "UserPoolAddOns",
	"email_configuration": "EmailConfiguration",
	"sms_configuration": "SmsConfiguration",
	"alias_attributes": "AliasAttributes",
	"enabled_mfas": "EnabledMfas",
	"email_verification_subject": "EmailVerificationSubject",
	"lambda_config": "LambdaConfig",
	"username_attributes": "UsernameAttributes",
	"auto_verified_attributes": "AutoVerifiedAttributes",
	"device_configuration": "DeviceConfiguration",
	"email_verification_message": "EmailVerificationMessage",
}

func ResourceCognitoUserPool() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolExists,
		Read: resourceCognitoUserPoolRead,
		Create: resourceCognitoUserPoolCreate,
		Update: resourceCognitoUserPoolUpdate,
		Delete: resourceCognitoUserPoolDelete,
		CustomizeDiff: resourceCognitoUserPoolCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"user_pool_tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyUserPoolPolicies(),
				Optional: true,
				MaxItems: 1,
			},
			"verification_message_template": {
				Type: schema.TypeList,
				Elem: propertyUserPoolVerificationMessageTemplate(),
				Optional: true,
				MaxItems: 1,
			},
			"mfa_configuration": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schema": {
				Type: schema.TypeList,
				Elem: propertyUserPoolSchemaAttribute(),
				Optional: true,
			},
			"admin_create_user_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolAdminCreateUserConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"sms_authentication_message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_pool_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"sms_verification_message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_pool_add_ons": {
				Type: schema.TypeList,
				Elem: propertyUserPoolUserPoolAddOns(),
				Optional: true,
				MaxItems: 1,
			},
			"email_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolEmailConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"sms_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolSmsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"alias_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled_mfas": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"email_verification_subject": {
				Type: schema.TypeString,
				Optional: true,
			},
			"lambda_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolLambdaConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"username_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"auto_verified_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"device_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolDeviceConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"email_verification_message": {
				Type: schema.TypeString,
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

func resourceCognitoUserPoolExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolType, ResourceCognitoUserPool(), data, meta)
}

func resourceCognitoUserPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolType, ResourceCognitoUserPool(), data, cognitoUserPoolProperties, meta)
}

func resourceCognitoUserPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolType, ResourceCognitoUserPool(), data, cognitoUserPoolProperties, meta)
}

func resourceCognitoUserPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolType, data, meta)
}

func resourceCognitoUserPoolCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolType, data, meta)
}
