// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var securityConfigurationEncryptionConfigurationProperties map[string]string = map[string]string{
	"s3_encryptions": "S3Encryptions",
	"cloud_watch_encryption": "CloudWatchEncryption",
	"job_bookmarks_encryption": "JobBookmarksEncryption",
}

func propertySecurityConfigurationEncryptionConfiguration(extras...string) *schema.Resource {
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
			"s3_encryptions": {
				Type: schema.TypeList,
				Elem: propertySecurityConfigurationS3Encryptions(),
				Optional: true,
				MaxItems: 1,
			},
			"cloud_watch_encryption": {
				Type: schema.TypeList,
				Elem: propertySecurityConfigurationCloudWatchEncryption(),
				Optional: true,
				MaxItems: 1,
			},
			"job_bookmarks_encryption": {
				Type: schema.TypeList,
				Elem: propertySecurityConfigurationJobBookmarksEncryption(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
