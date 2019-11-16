// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var distributionOriginProperties map[string]string = map[string]string{
	"origin_custom_headers": "OriginCustomHeaders",
	"domain_name": "DomainName",
	"s3_origin_config": "S3OriginConfig",
	"origin_path": "OriginPath",
	"id": "Id",
	"custom_origin_config": "CustomOriginConfig",
}

func propertyDistributionOrigin(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"origin_custom_headers": {
				Type: schema.TypeList,
				Elem: propertyDistributionOriginCustomHeader(),
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_origin_config": {
				Type: schema.TypeSet,
				Elem: propertyDistributionS3OriginConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"origin_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"custom_origin_config": {
				Type: schema.TypeSet,
				Elem: propertyDistributionCustomOriginConfig(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
