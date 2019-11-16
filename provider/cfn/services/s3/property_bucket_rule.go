// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html

package s3

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var bucketRuleProperties map[string]string = map[string]string{
	"abort_incomplete_multipart_upload": "AbortIncompleteMultipartUpload",
	"expiration_date": "ExpirationDate",
	"expiration_in_days": "ExpirationInDays",
	"id": "Id",
	"noncurrent_version_expiration_in_days": "NoncurrentVersionExpirationInDays",
	"noncurrent_version_transition": "NoncurrentVersionTransition",
	"noncurrent_version_transitions": "NoncurrentVersionTransitions",
	"prefix": "Prefix",
	"status": "Status",
	"tag_filters": "TagFilters",
	"transition": "Transition",
	"transitions": "Transitions",
}

func propertyBucketRule(extras...string) *schema.Resource {
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
			"abort_incomplete_multipart_upload": {
				Type: schema.TypeList,
				Elem: propertyBucketAbortIncompleteMultipartUpload(),
				Optional: true,
				MaxItems: 1,
			},
			"expiration_date": {
				Type: schema.TypeString,
				Optional: true,
			},
			"expiration_in_days": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"noncurrent_version_expiration_in_days": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"noncurrent_version_transition": {
				Type: schema.TypeList,
				Elem: propertyBucketNoncurrentVersionTransition(),
				Optional: true,
				MaxItems: 1,
			},
			"noncurrent_version_transitions": {
				Type: schema.TypeSet,
				Elem: propertyBucketNoncurrentVersionTransition(),
				Optional: true,
			},
			"prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"status": {
				Type: schema.TypeString,
				Required: true,
			},
			"tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyBucketTagFilter(),
				Optional: true,
			},
			"transition": {
				Type: schema.TypeList,
				Elem: propertyBucketTransition(),
				Optional: true,
				MaxItems: 1,
			},
			"transitions": {
				Type: schema.TypeSet,
				Elem: propertyBucketTransition(),
				Optional: true,
			},
		},
	}
}
