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

func propertyOrigin() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"origin_custom_headers": {
				Type: schema.TypeList,
				Elem: propertyOriginCustomHeader(),
				Required: false,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_origin_config": {
				Type: schema.TypeList,
				Elem: propertyS3OriginConfig(),
				Required: false,
				MaxItems: 1,
			},
			"origin_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"custom_origin_config": {
				Type: schema.TypeList,
				Elem: propertyCustomOriginConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}