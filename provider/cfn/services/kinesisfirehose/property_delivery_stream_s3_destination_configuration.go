// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamS3DestinationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"buffering_hints": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamBufferingHints(),
				Required: true,
				MaxItems: 1,
			},
			"cloud_watch_logging_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamCloudWatchLoggingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"compression_format": {
				Type: schema.TypeString,
				Required: true,
			},
			"encryption_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamEncryptionConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"error_output_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}