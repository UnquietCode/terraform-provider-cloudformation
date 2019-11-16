// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolRiskConfigurationAttachmentType string = "AWS::Cognito::UserPoolRiskConfigurationAttachment"

func ResourceCognitoUserPoolRiskConfigurationAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolRiskConfigurationAttachmentExists,
		Read: resourceCognitoUserPoolRiskConfigurationAttachmentRead,
		Create: resourceCognitoUserPoolRiskConfigurationAttachmentCreate,
		Update: resourceCognitoUserPoolRiskConfigurationAttachmentUpdate,
		Delete: resourceCognitoUserPoolRiskConfigurationAttachmentDelete,
		CustomizeDiff: resourceCognitoUserPoolRiskConfigurationAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"compromised_credentials_risk_configuration": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"account_takeover_risk_configuration": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"risk_exception_configuration": {
				Type: schema.TypeSet,
				Elem: propertyUserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCognitoUserPoolRiskConfigurationAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolRiskConfigurationAttachmentType, ResourceCognitoUserPoolRiskConfigurationAttachment(), data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoUserPoolRiskConfigurationAttachmentType, ResourceCognitoUserPoolRiskConfigurationAttachment(), data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoUserPoolRiskConfigurationAttachmentType, ResourceCognitoUserPoolRiskConfigurationAttachment(), data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoUserPoolRiskConfigurationAttachmentType, data, meta)
}

func resourceCognitoUserPoolRiskConfigurationAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoUserPoolRiskConfigurationAttachmentType, data, meta)
}
