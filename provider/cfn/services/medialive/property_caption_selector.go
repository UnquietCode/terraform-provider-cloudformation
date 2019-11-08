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

func propertyCaptionSelector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"language_code": {
				Type: schema.TypeString,
				Required: false,
			},
			"selector_settings": {
				Type: schema.TypeList,
				Elem: propertyCaptionSelectorSettings(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}