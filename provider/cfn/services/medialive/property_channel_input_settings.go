// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html

package medialive

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelInputSettings(extras...string) *schema.Resource {
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
			"deblock_filter": {
				Type: schema.TypeString,
				Optional: true,
			},
			"filter_strength": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"input_filter": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_end_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"video_selector": {
				Type: schema.TypeList,
				Elem: propertyChannelVideoSelector(),
				Optional: true,
				MaxItems: 1,
			},
			"audio_selectors": {
				Type: schema.TypeList,
				Elem: propertyChannelAudioSelector(),
				Optional: true,
			},
			"caption_selectors": {
				Type: schema.TypeList,
				Elem: propertyChannelCaptionSelector(),
				Optional: true,
			},
			"denoise_filter": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_input_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelNetworkInputSettings(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}