// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamExtendedS3DestinationConfiguration(extras...string) *schema.Resource {
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
			"data_format_conversion_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamDataFormatConversionConfiguration(),
				Required: false,
				MaxItems: 1,
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
			"processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamProcessingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_backup_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamS3DestinationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"s3_backup_mode": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}