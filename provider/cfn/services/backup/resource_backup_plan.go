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

func ResourceBackupPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupPlanCreate,
		Read:   resourceBackupPlanRead,
		Update: resourceBackupPlanUpdate,
		Delete: resourceBackupPlanDelete,

		Schema: map[string]*schema.Schema{
			"backup_plan": {
				Type: schema.TypeList,
				Elem: propertyBackupPlanBackupPlanResourceType(),
				Required: true,
				MaxItems: 1,
			},
			"backup_plan_tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceBackupPlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Backup::BackupPlan", data, meta)
}

func resourceBackupPlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Backup::BackupPlan", data, meta)
}

func resourceBackupPlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Backup::BackupPlan", data, meta)
}

func resourceBackupPlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Backup::BackupPlan", data, meta)
}