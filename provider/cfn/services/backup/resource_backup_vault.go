// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBackupVault() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupVaultCreate,
		Read:   resourceBackupVaultRead,
		Update: resourceBackupVaultUpdate,
		Delete: resourceBackupVaultDelete,

		Schema: map[string]*schema.Schema{
			"backup_vault_tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"backup_vault_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"encryption_key_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"notifications": {
				Type: schema.TypeList,
				Elem: propertyBackupVaultNotificationObjectType(),
				Required: false,
				MaxItems: 1,
			},
			"access_policy": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceBackupVaultCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupVaultRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupVaultUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupVaultDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Backup::BackupVault", data, meta)
}