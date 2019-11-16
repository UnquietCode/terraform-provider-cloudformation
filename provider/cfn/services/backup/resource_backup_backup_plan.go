// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupPlanType string = "AWS::Backup::BackupPlan"

var backupBackupPlanProperties map[string]string = map[string]string{
	"backup_plan": "BackupPlan",
	"backup_plan_tags": "BackupPlanTags",
}

func ResourceBackupBackupPlan() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBackupBackupPlanExists,
		Read: resourceBackupBackupPlanRead,
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
	return plugin.ResourceRead(backupBackupPlanType, ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(backupBackupPlanType, ResourceBackupBackupPlan(), data, backupBackupPlanProperties, meta)
}

func resourceBackupBackupPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(backupBackupPlanType, ResourceBackupBackupPlan(), data, backupBackupPlanProperties, meta)
}

func resourceBackupBackupPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(backupBackupPlanType, data, meta)
}

func resourceBackupBackupPlanCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(backupBackupPlanType, data, meta)
}
