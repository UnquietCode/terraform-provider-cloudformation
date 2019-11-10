// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBackupSelectionConditionResourceType() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"condition_value": {
				Type: schema.TypeString,
				Required: true,
			},
			"condition_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"condition_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}