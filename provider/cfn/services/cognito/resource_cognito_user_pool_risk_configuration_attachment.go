// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolRiskConfigurationAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceCognitoUserPoolRiskConfigurationAttachmentCreate,
		Read:   resourceCognitoUserPoolRiskConfigurationAttachmentRead,
		Update: resourceCognitoUserPoolRiskConfigurationAttachmentUpdate,
		Delete: resourceCognitoUserPoolRiskConfigurationAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"compromised_credentials_risk_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType(),
				Required: false,
				MaxItems: 1,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_takeover_risk_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType(),
				Required: false,
				MaxItems: 1,
			},
			"risk_exception_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceCognitoUserPoolRiskConfigurationAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolRiskConfigurationAttachment", data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolRiskConfigurationAttachment", data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolRiskConfigurationAttachment", data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolRiskConfigurationAttachment", data, meta)
}