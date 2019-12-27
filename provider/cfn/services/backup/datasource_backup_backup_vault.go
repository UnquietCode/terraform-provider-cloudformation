// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html

package backup

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupVaultType string = "AWS::Backup::BackupVault"

func DatasourceBackupBackupVault() *schema.Resource {
	return &schema.Resource{
		Read: datasourceBackupBackupVaultRead,
		
		Schema: map[string]*schema.Schema{
			"backup_vault_tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				Type: schema.TypeList,
				Elem: propertyBackupVaultNotificationObjectType(),
				Optional: true,
				MaxItems: 1,
			},
			"access_policy": {
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceBackupBackupVaultRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(backupBackupVaultType, DatasourceBackupBackupVault(), data, meta)
}
