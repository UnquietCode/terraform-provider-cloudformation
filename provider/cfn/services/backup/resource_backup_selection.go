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

func ResourceBackupSelection() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupSelectionCreate,
		Read:   resourceBackupSelectionRead,
		Update: resourceBackupSelectionUpdate,
		Delete: resourceBackupSelectionDelete,

		Schema: map[string]*schema.Schema{
			"backup_selection": {
				Type: schema.TypeList,
				Elem: propertyBackupSelectionBackupSelectionResourceType(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"backup_plan_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBackupSelectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Backup::BackupSelection", data, meta)
}

func resourceBackupSelectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Backup::BackupSelection", data, meta)
}

func resourceBackupSelectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Backup::BackupSelection", data, meta)
}

func resourceBackupSelectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Backup::BackupSelection", data, meta)
}