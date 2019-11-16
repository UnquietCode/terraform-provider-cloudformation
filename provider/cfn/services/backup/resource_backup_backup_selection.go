// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupSelectionType string = "AWS::Backup::BackupSelection"

func ResourceBackupBackupSelection() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBackupBackupSelectionExists,
		Read: resourceBackupBackupSelectionRead,
		Create: resourceBackupBackupSelectionCreate,
		Update: resourceBackupBackupSelectionUpdate,
		Delete: resourceBackupBackupSelectionDelete,
		CustomizeDiff: resourceBackupBackupSelectionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"backup_selection": {
				Type: schema.TypeSet,
				Elem: propertyBackupSelectionBackupSelectionResourceType(),
				Required: true,
				MaxItems: 1,
			},
			"backup_plan_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBackupBackupSelectionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceBackupBackupSelectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(backupBackupSelectionType, ResourceBackupBackupSelection(), data, meta)
}

func resourceBackupBackupSelectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(backupBackupSelectionType, ResourceBackupBackupSelection(), data, meta)
}

func resourceBackupBackupSelectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(backupBackupSelectionType, ResourceBackupBackupSelection(), data, meta)
}

func resourceBackupBackupSelectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(backupBackupSelectionType, data, meta)
}

func resourceBackupBackupSelectionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(backupBackupSelectionType, data, meta)
}
