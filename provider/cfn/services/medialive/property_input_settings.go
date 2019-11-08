// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInputSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deblock_filter": {
				Type: schema.TypeString,
				Required: false,
			},
			"filter_strength": {
				Type: schema.TypeInt,
				Required: false,
			},
			"input_filter": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_end_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"video_selector": {
				Type: schema.TypeList,
				Elem: propertyVideoSelector(),
				Required: false,
				MaxItems: 1,
			},
			"audio_selectors": {
				Type: schema.TypeList,
				Elem: propertyAudioSelector(),
				Required: false,
			},
			"caption_selectors": {
				Type: schema.TypeList,
				Elem: propertyCaptionSelector(),
				Required: false,
			},
			"denoise_filter": {
				Type: schema.TypeString,
				Required: false,
			},
			"network_input_settings": {
				Type: schema.TypeList,
				Elem: propertyNetworkInputSettings(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}