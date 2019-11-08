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

func propertyCaptionSelectorSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dvb_sub_source_settings": {
				Type: schema.TypeList,
				Elem: propertyDvbSubSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"scte27_source_settings": {
				Type: schema.TypeList,
				Elem: propertyScte27SourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"arib_source_settings": {
				Type: schema.TypeList,
				Elem: propertyAribSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"embedded_source_settings": {
				Type: schema.TypeList,
				Elem: propertyEmbeddedSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"scte20_source_settings": {
				Type: schema.TypeList,
				Elem: propertyScte20SourceSettings(),
				Required: false,
				MaxItems: 1,
			},
			"teletext_source_settings": {
				Type: schema.TypeList,
				Elem: propertyTeletextSourceSettings(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}