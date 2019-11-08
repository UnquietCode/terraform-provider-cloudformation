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

func propertyInputAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"input_attachment_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"input_settings": {
				Type: schema.TypeList,
				Elem: propertyInputSettings(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}