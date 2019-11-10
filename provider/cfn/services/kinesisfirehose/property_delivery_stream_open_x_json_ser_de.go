// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamOpenXJsonSerDe() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"case_insensitive": {
				Type: schema.TypeBool,
				Required: false,
			},
			"column_to_json_key_mappings": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"convert_dots_in_json_keys_to_underscores": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}