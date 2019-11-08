// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCsvClassifier() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"quote_symbol": {
				Type: schema.TypeString,
				Required: false,
			},
			"contains_header": {
				Type: schema.TypeString,
				Required: false,
			},
			"delimiter": {
				Type: schema.TypeString,
				Required: false,
			},
			"header": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"allow_single_column": {
				Type: schema.TypeBool,
				Required: false,
			},
			"disable_value_trimming": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}