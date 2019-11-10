// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html

package s3

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketReplicationDestination(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_control_translation": {
				Type: schema.TypeList,
				Elem: propertyBucketAccessControlTranslation(),
				Required: false,
				MaxItems: 1,
			},
			"account": {
				Type: schema.TypeString,
				Required: false,
			},
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"encryption_configuration": {
				Type: schema.TypeList,
				Elem: propertyBucketEncryptionConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"storage_class": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}