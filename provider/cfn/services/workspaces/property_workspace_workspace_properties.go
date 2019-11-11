// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html

package workspaces

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWorkspaceWorkspaceProperties(extras...string) *schema.Resource {
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
			"compute_type_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"root_volume_size_gib": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"running_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"running_mode_auto_stop_timeout_in_minutes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"user_volume_size_gib": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}

