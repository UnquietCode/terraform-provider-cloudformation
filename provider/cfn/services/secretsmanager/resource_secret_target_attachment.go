// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package secretsmanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecretTargetAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretTargetAttachmentCreate,
		Read:   resourceSecretTargetAttachmentRead,
		Update: resourceSecretTargetAttachmentUpdate,
		Delete: resourceSecretTargetAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSecretTargetAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecretsManager::SecretTargetAttachment", data, meta)
}

func resourceSecretTargetAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecretsManager::SecretTargetAttachment", data, meta)
}

func resourceSecretTargetAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecretsManager::SecretTargetAttachment", data, meta)
}

func resourceSecretTargetAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecretsManager::SecretTargetAttachment", data, meta)
}