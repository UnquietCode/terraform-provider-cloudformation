// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html

package backup

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func propertyBackupPlanBackupRuleResourceType(extras...string) *schema.Resource {
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
			"completion_window_minutes": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"recovery_point_tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"lifecycle": {
				Type: schema.TypeList,
				Elem: propertyBackupPlanLifecycleResourceType(),
				Optional: true,
				MaxItems: 1,
			},
			"target_backup_vault": {
				Type: schema.TypeString,
				Required: true,
			},
			"start_window_minutes": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"rule_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
