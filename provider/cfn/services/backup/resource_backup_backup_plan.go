// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBackupBackupPlan() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBackupBackupPlanExists,
		Read:   resourceBackupBackupPlanRead,
		Create: resourceBackupBackupPlanCreate,
		Update: resourceBackupBackupPlanUpdate,
		Delete: resourceBackupBackupPlanDelete,
		CustomizeDiff: resourceBackupBackupPlanCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"backup_plan": {
				Type: schema.TypeList,
				Elem: propertyBackupPlanBackupPlanResourceType(),
				Required: true,
				MaxItems: 1,
			},
			"backup_plan_tags": {
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

func resourceBackupBackupPlanExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceBackupBackupPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Backup::BackupPlan", ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Backup::BackupPlan", ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Backup::BackupPlan", ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Backup::BackupPlan", data, meta)
}

func resourceBackupBackupPlanCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}