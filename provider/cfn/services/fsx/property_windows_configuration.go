// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package fsx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWindowsConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"self_managed_active_directory_configuration": {
				Type: schema.TypeList,
				Elem: propertySelfManagedActiveDirectoryConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"weekly_maintenance_start_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"active_directory_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"throughput_capacity": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"copy_tags_to_backups": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"daily_automatic_backup_start_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"automatic_backup_retention_days": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}