// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelVideoSelector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"selector_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelVideoSelectorSettings(),
				Required: false,
				MaxItems: 1,
			},
			"color_space": {
				Type: schema.TypeString,
				Required: false,
			},
			"color_space_usage": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}