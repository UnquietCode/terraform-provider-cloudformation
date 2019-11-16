// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoIdentityPoolRoleAttachmentType string = "AWS::Cognito::IdentityPoolRoleAttachment"

var cognitoIdentityPoolRoleAttachmentProperties map[string]string = map[string]string{
	"role_mappings": "RoleMappings",
	"identity_pool_id": "IdentityPoolId",
	"roles": "Roles",
}

func ResourceCognitoIdentityPoolRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCognitoIdentityPoolRoleAttachmentExists,
		Read: resourceCognitoIdentityPoolRoleAttachmentRead,
		Create: resourceCognitoIdentityPoolRoleAttachmentCreate,
		Update: resourceCognitoIdentityPoolRoleAttachmentUpdate,
		Delete: resourceCognitoIdentityPoolRoleAttachmentDelete,
		CustomizeDiff: resourceCognitoIdentityPoolRoleAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"role_mappings": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"identity_pool_id": {
				Type: schema.TypeString,
				Required: true,
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

func resourceCognitoIdentityPoolRoleAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, cognitoIdentityPoolRoleAttachmentProperties, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, cognitoIdentityPoolRoleAttachmentProperties, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoIdentityPoolRoleAttachmentType, data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoIdentityPoolRoleAttachmentType, data, meta)
}
