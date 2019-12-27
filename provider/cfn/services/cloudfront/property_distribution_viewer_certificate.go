// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html

package cloudfront

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionViewerCertificate(extras...string) *schema.Resource {
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
			"iam_certificate_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssl_support_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"minimum_protocol_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cloud_front_default_certificate": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"acm_certificate_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
