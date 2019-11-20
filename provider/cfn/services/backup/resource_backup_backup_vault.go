// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupVaultType string = "AWS::Backup::BackupVault"

func ResourceBackupBackupVault() *schema.Resource {
	return &schema.Resource{
		Read: resourceBackupBackupVaultRead,
		Create: resourceBackupBackupVaultCreate,
		Update: resourceBackupBackupVaultUpdate,
		Delete: resourceBackupBackupVaultDelete,
		CustomizeDiff: resourceBackupBackupVaultCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"backup_vault_tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"backup_vault_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"encryption_key_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notifications": {
				Type: schema.TypeSet,
				Elem: propertyBackupVaultNotificationObjectType(),
				Optional: true,
				MaxItems: 1,
			},
			"access_policy": {
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

func resourceBackupBackupVaultRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(backupBackupVaultType, ResourceBackupBackupVault(), data, meta)
}

func resourceBackupBackupVaultCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(backupBackupVaultType, ResourceBackupBackupVault(), data, meta)
}

func resourceBackupBackupVaultUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(backupBackupVaultType, ResourceBackupBackupVault(), data, meta)
}

func resourceBackupBackupVaultDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(backupBackupVaultType, data, meta)
}

func resourceBackupBackupVaultCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(backupBackupVaultType, data, meta)
}
