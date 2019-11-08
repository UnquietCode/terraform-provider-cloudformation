// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyForwardedValues() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookies": {
				Type: schema.TypeList,
				Elem: propertyCookies(),
				Required: false,
				MaxItems: 1,
			},
			"headers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"query_string": {
				Type: schema.TypeBool,
				Required: true,
			},
			"query_string_cache_keys": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}