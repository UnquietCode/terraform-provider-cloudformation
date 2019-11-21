// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html

package backup

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupPlanType string = "AWS::Backup::BackupPlan"

func ResourceBackupBackupPlan() *schema.Resource {
	return &schema.Resource{
		Read: resourceBackupBackupPlanRead,
		Create: resourceBackupBackupPlanCreate,
		Update: resourceBackupBackupPlanUpdate,
		Delete: resourceBackupBackupPlanDelete,
		CustomizeDiff: resourceBackupBackupPlanCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"backup_plan": {
				Type: schema.TypeSet,
				Elem: propertyBackupPlanBackupPlanResourceType(),
				Required: true,
				MaxItems: 1,
			},
			"backup_plan_tags": {
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

func resourceBackupBackupPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(backupBackupPlanType, ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(backupBackupPlanType, ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(backupBackupPlanType, ResourceBackupBackupPlan(), data, meta)
}

func resourceBackupBackupPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(backupBackupPlanType, data, meta)
}

func resourceBackupBackupPlanCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(backupBackupPlanType, data, meta)
}
