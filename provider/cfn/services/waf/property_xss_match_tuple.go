// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyXssMatchTuple() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_to_match": {
				Type: schema.TypeList,
				Elem: propertyFieldToMatch(),
				Required: true,
				MaxItems: 1,
			},
			"text_transformation": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}