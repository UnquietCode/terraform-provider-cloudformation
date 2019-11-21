// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html

package cognito

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cognitoIdentityPoolRoleAttachmentType string = "AWS::Cognito::IdentityPoolRoleAttachment"

func ResourceCognitoIdentityPoolRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Read: resourceCognitoIdentityPoolRoleAttachmentRead,
		Create: resourceCognitoIdentityPoolRoleAttachmentCreate,
		Update: resourceCognitoIdentityPoolRoleAttachmentUpdate,
		Delete: resourceCognitoIdentityPoolRoleAttachmentDelete,
		CustomizeDiff: resourceCognitoIdentityPoolRoleAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"role_mappings": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"identity_pool_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"roles": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCognitoIdentityPoolRoleAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cognitoIdentityPoolRoleAttachmentType, ResourceCognitoIdentityPoolRoleAttachment(), data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cognitoIdentityPoolRoleAttachmentType, data, meta)
}

func resourceCognitoIdentityPoolRoleAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cognitoIdentityPoolRoleAttachmentType, data, meta)
}
