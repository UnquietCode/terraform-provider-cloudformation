// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html

package fsx

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var fileSystemWindowsConfigurationProperties map[string]string = map[string]string{
	"self_managed_active_directory_configuration": "SelfManagedActiveDirectoryConfiguration",
	"weekly_maintenance_start_time": "WeeklyMaintenanceStartTime",
	"active_directory_id": "ActiveDirectoryId",
	"throughput_capacity": "ThroughputCapacity",
	"copy_tags_to_backups": "CopyTagsToBackups",
	"daily_automatic_backup_start_time": "DailyAutomaticBackupStartTime",
	"automatic_backup_retention_days": "AutomaticBackupRetentionDays",
}

func propertyFileSystemWindowsConfiguration(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"self_managed_active_directory_configuration": {
				Type: schema.TypeList,
				Elem: propertyFileSystemSelfManagedActiveDirectoryConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"weekly_maintenance_start_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"active_directory_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"throughput_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"copy_tags_to_backups": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"daily_automatic_backup_start_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"automatic_backup_retention_days": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
