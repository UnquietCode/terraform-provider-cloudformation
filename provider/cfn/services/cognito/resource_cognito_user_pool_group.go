// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolGroupExists,
		Read:   resourceCognitoUserPoolGroupRead,
		Create: resourceCognitoUserPoolGroupCreate,
		Update: resourceCognitoUserPoolGroupUpdate,
		Delete: resourceCognitoUserPoolGroupDelete,
		CustomizeDiff: resourceCognitoUserPoolGroupCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"precedence": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"role_arn": {
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

func resourceCognitoUserPoolGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolGroup", ResourceCognitoUserPoolGroup(), data, meta)
}

func resourceCognitoUserPoolGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolGroup", ResourceCognitoUserPoolGroup(), data, meta)
}

func resourceCognitoUserPoolGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolGroup", ResourceCognitoUserPoolGroup(), data, meta)
}

func resourceCognitoUserPoolGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolGroup", data, meta)
}

func resourceCognitoUserPoolGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}