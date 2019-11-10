// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelVideoSelectorSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"video_selector_program_id": {
				Type: schema.TypeList,
				Elem: propertyChannelVideoSelectorProgramId(),
				Required: false,
				MaxItems: 1,
			},
			"video_selector_pid": {
				Type: schema.TypeList,
				Elem: propertyChannelVideoSelectorPid(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}