// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var distributionCustomErrorResponseProperties map[string]string = map[string]string{
	"response_code": "ResponseCode",
	"error_caching_min_ttl": "ErrorCachingMinTTL",
	"error_code": "ErrorCode",
	"response_page_path": "ResponsePagePath",
}

func propertyDistributionCustomErrorResponse(extras...string) *schema.Resource {
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
			"response_code": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"error_caching_min_ttl": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"error_code": {
				Type: schema.TypeInt,
				Required: true,
			},
			"response_page_path": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
