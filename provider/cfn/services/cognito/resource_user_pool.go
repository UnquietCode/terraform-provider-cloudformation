// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceUserPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPoolCreate,
		Read:   resourceUserPoolRead,
		Update: resourceUserPoolUpdate,
		Delete: resourceUserPoolDelete,

		Schema: map[string]*schema.Schema{
			"user_pool_tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyUserPoolPolicies(),
				Required: false,
				MaxItems: 1,
			},
			"verification_message_template": {
				Type: schema.TypeList,
				Elem: propertyUserPoolVerificationMessageTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"mfa_configuration": {
				Type: schema.TypeString,
				Required: false,
			},
			"schema": {
				Type: schema.TypeList,
				Elem: propertyUserPoolSchemaAttribute(),
				Required: false,
			},
			"admin_create_user_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolAdminCreateUserConfig(),
				Required: false,
				MaxItems: 1,
			},
			"sms_authentication_message": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_pool_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"sms_verification_message": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_pool_add_ons": {
				Type: schema.TypeList,
				Elem: propertyUserPoolUserPoolAddOns(),
				Required: false,
				MaxItems: 1,
			},
			"email_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolEmailConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"sms_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolSmsConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"alias_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"enabled_mfas": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"email_verification_subject": {
				Type: schema.TypeString,
				Required: false,
			},
			"lambda_config": {
				Type: schema.TypeList,
				Elem: propertyUserPoolLambdaConfig(),
				Required: false,
				MaxItems: 1,
			},
			"username_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"auto_verified_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"device_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolDeviceConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"email_verification_message": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceUserPoolCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPool", data, meta)
}

func resourceUserPoolRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPool", data, meta)
}

func resourceUserPoolUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPool", data, meta)
}

func resourceUserPoolDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPool", data, meta)
}