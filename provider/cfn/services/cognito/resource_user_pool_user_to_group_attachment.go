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

func ResourceUserPoolUserToGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPoolUserToGroupAttachmentCreate,
		Read:   resourceUserPoolUserToGroupAttachmentRead,
		Update: resourceUserPoolUserToGroupAttachmentUpdate,
		Delete: resourceUserPoolUserToGroupAttachmentDelete,

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
		},
	}
}

func resourceUserPoolUserToGroupAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}

func resourceUserPoolUserToGroupAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}

func resourceUserPoolUserToGroupAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}

func resourceUserPoolUserToGroupAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolUserToGroupAttachment", data, meta)
}