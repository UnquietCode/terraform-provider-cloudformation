// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyByteMatchSetByteMatchTuple() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"target_string": {
				Type: schema.TypeString,
				Required: false,
			},
			"target_string_base64": {
				Type: schema.TypeString,
				Required: false,
			},
			"positional_constraint": {
				Type: schema.TypeString,
				Required: true,
			},
			"text_transformation": {
				Type: schema.TypeString,
				Required: true,
			},
			"field_to_match": {
				Type: schema.TypeList,
				Elem: propertyByteMatchSetFieldToMatch(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}