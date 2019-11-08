// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package workspaces

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWorkspaceProperties() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compute_type_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"root_volume_size_gib": {
				Type: schema.TypeInt,
				Required: false,
			},
			"running_mode": {
				Type: schema.TypeString,
				Required: false,
			},
			"running_mode_auto_stop_timeout_in_minutes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"user_volume_size_gib": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}