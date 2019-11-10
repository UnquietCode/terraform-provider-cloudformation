// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBackupBackupVault() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupBackupVaultCreate,
		Read:   resourceBackupBackupVaultRead,
		Update: resourceBackupBackupVaultUpdate,
		Delete: resourceBackupBackupVaultDelete,

		Schema: map[string]*schema.Schema{
			"backup_vault_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backup_vault_arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"backup_vault_tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"encryption_key_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"notifications": {
				Type: schema.TypeList,
				Elem: propertyBackupVaultNotificationObjectType(),
				Optional: true,
				MaxItems: 1,
			},
			"access_policy": {
				Type: schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceBackupBackupVaultCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupBackupVaultRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupBackupVaultUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Backup::BackupVault", data, meta)
}

func resourceBackupBackupVaultDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Backup::BackupVault", data, meta)
}