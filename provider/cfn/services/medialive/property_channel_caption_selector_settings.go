// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html

package medialive

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var channelCaptionSelectorSettingsProperties map[string]string = map[string]string{
	"dvb_sub_source_settings": "DvbSubSourceSettings",
	"scte27_source_settings": "Scte27SourceSettings",
	"arib_source_settings": "AribSourceSettings",
	"embedded_source_settings": "EmbeddedSourceSettings",
	"scte20_source_settings": "Scte20SourceSettings",
	"teletext_source_settings": "TeletextSourceSettings",
}

func propertyChannelCaptionSelectorSettings(extras...string) *schema.Resource {
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
			"dvb_sub_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelDvbSubSourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"scte27_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelScte27SourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"arib_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelAribSourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"embedded_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelEmbeddedSourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"scte20_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelScte20SourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"teletext_source_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelTeletextSourceSettings(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
