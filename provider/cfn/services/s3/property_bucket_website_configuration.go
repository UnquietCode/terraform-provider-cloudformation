// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html

package s3

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketWebsiteConfiguration(extras...string) *schema.Resource {
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
			"error_document": {
				Type: schema.TypeString,
				Optional: true,
			},
			"index_document": {
				Type: schema.TypeString,
				Optional: true,
			},
			"redirect_all_requests_to": {
				Type: schema.TypeSet,
				Elem: propertyBucketRedirectAllRequestsTo(),
				Optional: true,
				MaxItems: 1,
			},
			"routing_rules": {
				Type: schema.TypeSet,
				Elem: propertyBucketRoutingRule(),
				Optional: true,
			},
		},
	}
}
