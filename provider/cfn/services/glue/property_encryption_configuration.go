// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEncryptionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"s3_encryptions": {
				Type: schema.TypeList,
				Elem: propertyS3Encryptions(),
				Required: false,
				MaxItems: 1,
			},
			"cloud_watch_encryption": {
				Type: schema.TypeList,
				Elem: propertyCloudWatchEncryption(),
				Required: false,
				MaxItems: 1,
			},
			"job_bookmarks_encryption": {
				Type: schema.TypeList,
				Elem: propertyJobBookmarksEncryption(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}