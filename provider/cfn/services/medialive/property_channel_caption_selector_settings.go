// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package medialive

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelCaptionSelectorSettings(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dvb_sub_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelDvbSubSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"scte27_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelScte27SourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"arib_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelAribSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"embedded_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelEmbeddedSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"scte20_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelScte20SourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"teletext_source_settings": {
				Type: schema.TypeList,
				Elem: propertyChannelTeletextSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}