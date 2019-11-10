// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMemberApprovalThresholdPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"threshold_comparator": {
				Type: schema.TypeString,
				Required: false,
			},
			"threshold_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
			"proposal_duration_in_hours": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}