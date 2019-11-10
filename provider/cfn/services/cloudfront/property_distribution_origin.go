// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionOrigin() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"origin_custom_headers": {
				Type: schema.TypeList,
				Elem: propertyDistributionOriginCustomHeader(),
				Required: false,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_origin_config": {
				Type: schema.TypeList,
				Elem: propertyDistributionS3OriginConfig(),
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
				Elem: propertyDistributionCustomOriginConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}