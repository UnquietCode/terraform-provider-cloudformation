// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-networkinputsettings.html

package medialive

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var channelNetworkInputSettingsProperties map[string]string = map[string]string{
	"server_validation": "ServerValidation",
	"hls_input_settings": "HlsInputSettings",
}

func propertyChannelNetworkInputSettings(extras...string) *schema.Resource {
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
			"server_validation": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hls_input_settings": {
				Type: schema.TypeSet,
				Elem: propertyChannelHlsInputSettings(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
