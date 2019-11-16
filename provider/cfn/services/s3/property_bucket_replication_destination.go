// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html

package s3

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var bucketReplicationDestinationProperties map[string]string = map[string]string{
	"access_control_translation": "AccessControlTranslation",
	"account": "Account",
	"bucket": "Bucket",
	"encryption_configuration": "EncryptionConfiguration",
	"storage_class": "StorageClass",
}

func propertyBucketReplicationDestination(extras...string) *schema.Resource {
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
			"access_control_translation": {
				Type: schema.TypeSet,
				Elem: propertyBucketAccessControlTranslation(),
				Optional: true,
				MaxItems: 1,
			},
			"account": {
				Type: schema.TypeString,
				Optional: true,
			},
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"encryption_configuration": {
				Type: schema.TypeSet,
				Elem: propertyBucketEncryptionConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"storage_class": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
