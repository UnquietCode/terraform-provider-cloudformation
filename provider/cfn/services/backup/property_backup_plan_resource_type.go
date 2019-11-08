// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBackupPlanResourceType() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_plan_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"backup_plan_rule": {
				Type: schema.TypeList,
				Elem: propertyBackupRuleResourceType(),
				Required: true,
			},
		},
	}
}