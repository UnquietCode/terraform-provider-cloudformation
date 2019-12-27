// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html

package backup

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const backupBackupSelectionType string = "AWS::Backup::BackupSelection"

func DatasourceBackupBackupSelection() *schema.Resource {
	return &schema.Resource{
		Read: datasourceBackupBackupSelectionRead,
		
		Schema: map[string]*schema.Schema{
			"backup_selection": {
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceBackupBackupSelectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(backupBackupSelectionType, DatasourceBackupBackupSelection(), data, meta)
}
