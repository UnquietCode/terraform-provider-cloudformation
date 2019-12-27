// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolRiskConfigurationAttachmentType string = "AWS::Cognito::UserPoolRiskConfigurationAttachment"

func DatasourceCognitoUserPoolRiskConfigurationAttachment() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCognitoUserPoolRiskConfigurationAttachmentRead,
		
		Schema: map[string]*schema.Schema{
			"compromised_credentials_risk_configuration": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"risk_exception_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCognitoUserPoolRiskConfigurationAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolRiskConfigurationAttachmentType, DatasourceCognitoUserPoolRiskConfigurationAttachment(), data, meta)
}
