// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceCognitoUserPoolUserToGroupAttachmentCreate,
		Read:   resourceCognitoUserPoolUserToGroupAttachmentRead,
		Delete: resourceCognitoUserPoolUserToGroupAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCognitoUserPoolUserToGroupAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolUserToGroupAttachment", ResourceCognitoUserPoolUserToGroupAttachment(), data, meta)
}

func resourceCognitoUserPoolUserToGroupAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}