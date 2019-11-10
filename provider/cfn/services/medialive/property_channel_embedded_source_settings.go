// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelEmbeddedSourceSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source608_channel_number": {
				Type: schema.TypeInt,
				Required: false,
			},
			"scte20_detection": {
				Type: schema.TypeString,
				Required: false,
			},
			"source608_track_number": {
				Type: schema.TypeInt,
				Required: false,
			},
			"convert608_to708": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}