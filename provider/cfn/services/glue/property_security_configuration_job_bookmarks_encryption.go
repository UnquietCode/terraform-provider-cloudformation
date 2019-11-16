// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-jobbookmarksencryption.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var securityConfigurationJobBookmarksEncryptionProperties map[string]string = map[string]string{
	"kms_key_arn": "KmsKeyArn",
	"job_bookmarks_encryption_mode": "JobBookmarksEncryptionMode",
}

func propertySecurityConfigurationJobBookmarksEncryption(extras...string) *schema.Resource {
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
			"kms_key_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"job_bookmarks_encryption_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
