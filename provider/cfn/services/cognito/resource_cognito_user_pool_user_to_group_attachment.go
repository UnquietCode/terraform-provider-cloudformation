// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoUserPoolUserToGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoUserPoolUserToGroupAttachmentExists,
		Read: resourceCognitoUserPoolUserToGroupAttachmentRead,
		Create: resourceCognitoUserPoolUserToGroupAttachmentCreate,
		Update: resourceCognitoUserPoolUserToGroupAttachmentUpdate,
		Delete: resourceCognitoUserPoolUserToGroupAttachmentDelete,
		CustomizeDiff: resourceCognitoUserPoolUserToGroupAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"username": {
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

func resourceCognitoUserPoolUserToGroupAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}

