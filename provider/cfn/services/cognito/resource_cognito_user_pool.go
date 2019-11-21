// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolType string = "AWS::Cognito::UserPool"

func ResourceCognitoUserPool() *schema.Resource {
	return &schema.Resource{
		Read: resourceCognitoUserPoolRead,
		Create: resourceCognitoUserPoolCreate,
		Update: resourceCognitoUserPoolUpdate,
		Delete: resourceCognitoUserPoolDelete,
		CustomizeDiff: resourceCognitoUserPoolCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"user_pool_tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"policies": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolPolicies(),
				Optional: true,
				MaxItems: 1,
			},
			"verification_message_template": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyUserPoolUserPoolAddOns(),
				Optional: true,
				MaxItems: 1,
			},
			"email_configuration": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolEmailConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"sms_configuration": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCognitoUserPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolType, ResourceCognitoUserPool(), data, meta)
}

func resourceCognitoUserPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolType, ResourceCognitoUserPool(), data, meta)
}

func resourceCognitoUserPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolType, ResourceCognitoUserPool(), data, meta)
}

func resourceCognitoUserPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolType, data, meta)
}

func resourceCognitoUserPoolCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolType, data, meta)
}
