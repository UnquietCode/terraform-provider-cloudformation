// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"abort_incomplete_multipart_upload": {
				Type: schema.TypeList,
				Elem: propertyAbortIncompleteMultipartUpload(),
				Required: false,
				MaxItems: 1,
			},
			"expiration_date": {
				Type: schema.TypeString,
				Required: false,
			},
			"expiration_in_days": {
				Type: schema.TypeInt,
				Required: false,
			},
			"id": {
				Type: schema.TypeString,
				Required: false,
			},
			"noncurrent_version_expiration_in_days": {
				Type: schema.TypeInt,
				Required: false,
			},
			"noncurrent_version_transition": {
				Type: schema.TypeList,
				Elem: propertyNoncurrentVersionTransition(),
				Required: false,
				MaxItems: 1,
			},
			"noncurrent_version_transitions": {
				Type: schema.TypeSet,
				Elem: propertyNoncurrentVersionTransition(),
				Required: false,
			},
			"prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"status": {
				Type: schema.TypeString,
				Required: true,
			},
			"tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyTagFilter(),
				Required: false,
			},
			"transition": {
				Type: schema.TypeList,
				Elem: propertyTransition(),
				Required: false,
				MaxItems: 1,
			},
			"transitions": {
				Type: schema.TypeSet,
				Elem: propertyTransition(),
				Required: false,
			},
		},
	}
}