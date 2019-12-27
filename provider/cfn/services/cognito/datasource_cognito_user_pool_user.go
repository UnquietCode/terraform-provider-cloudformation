// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoUserPoolUserType string = "AWS::Cognito::UserPoolUser"

func DatasourceCognitoUserPoolUser() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCognitoUserPoolUserRead,
		
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
			"client_metadata": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCognitoUserPoolUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoUserPoolUserType, DatasourceCognitoUserPoolUser(), data, meta)
}
