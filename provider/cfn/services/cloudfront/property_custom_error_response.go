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

func propertyCustomErrorResponse() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"response_code": {
				Type: schema.TypeInt,
				Required: false,
			},
			"error_caching_min_ttl": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"error_code": {
				Type: schema.TypeInt,
				Required: true,
			},
			"response_page_path": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}