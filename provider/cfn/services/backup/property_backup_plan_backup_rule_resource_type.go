// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package backup

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBackupPlanBackupRuleResourceType(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"completion_window_minutes": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"recovery_point_tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"lifecycle": {
				Type: schema.TypeList,
				Elem: propertyBackupPlanLifecycleResourceType(),
				Required: false,
				MaxItems: 1,
			},
			"target_backup_vault": {
				Type: schema.TypeString,
				Required: true,
			},
			"start_window_minutes": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"rule_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}