// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCognitoIdentityPoolRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceCognitoIdentityPoolRoleAttachmentCreate,
		Read:   resourceCognitoIdentityPoolRoleAttachmentRead,
		Update: resourceCognitoIdentityPoolRoleAttachmentUpdate,
		Delete: resourceCognitoIdentityPoolRoleAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"role_mappings": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"identity_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"roles": {
				Type: schema.TypeMap,
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

func resourceCognitoIdentityPoolRoleAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::IdentityPoolRoleAttachment", ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::IdentityPoolRoleAttachment", ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::IdentityPoolRoleAttachment", ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::IdentityPoolRoleAttachment", data, meta)
}