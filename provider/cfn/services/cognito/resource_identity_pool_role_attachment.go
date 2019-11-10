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

func ResourceIdentityPoolRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPoolRoleAttachmentCreate,
		Read:   resourceIdentityPoolRoleAttachmentRead,
		Update: resourceIdentityPoolRoleAttachmentUpdate,
		Delete: resourceIdentityPoolRoleAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"role_mappings": {
				Type: schema.TypeMap,
				Required: false,
			},
			"identity_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"roles": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceIdentityPoolRoleAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::IdentityPoolRoleAttachment", data, meta)
}

func resourceIdentityPoolRoleAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::IdentityPoolRoleAttachment", data, meta)
}

func resourceIdentityPoolRoleAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::IdentityPoolRoleAttachment", data, meta)
}

func resourceIdentityPoolRoleAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::IdentityPoolRoleAttachment", data, meta)
}