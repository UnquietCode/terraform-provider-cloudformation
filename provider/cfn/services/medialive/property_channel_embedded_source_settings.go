// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html

package medialive

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var channelEmbeddedSourceSettingsProperties map[string]string = map[string]string{
	"source608_channel_number": "Source608ChannelNumber",
	"scte20_detection": "Scte20Detection",
	"source608_track_number": "Source608TrackNumber",
	"convert608_to708": "Convert608To708",
}

func propertyChannelEmbeddedSourceSettings(extras...string) *schema.Resource {
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
			"source608_channel_number": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"scte20_detection": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source608_track_number": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"convert608_to708": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
