// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInputTransformer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"input_paths_map": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"input_template": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}