// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkdestinationconfiguration.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var deliveryStreamSplunkDestinationConfigurationProperties map[string]string = map[string]string{
	"cloud_watch_logging_options": "CloudWatchLoggingOptions",
	"hec_acknowledgment_timeout_in_seconds": "HECAcknowledgmentTimeoutInSeconds",
	"hec_endpoint": "HECEndpoint",
	"hec_endpoint_type": "HECEndpointType",
	"hec_token": "HECToken",
	"processing_configuration": "ProcessingConfiguration",
	"retry_options": "RetryOptions",
	"s3_backup_mode": "S3BackupMode",
	"s3_configuration": "S3Configuration",
}

func propertyDeliveryStreamSplunkDestinationConfiguration(extras...string) *schema.Resource {
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
			"cloud_watch_logging_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamCloudWatchLoggingOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"hec_acknowledgment_timeout_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"hec_endpoint": {
				Type: schema.TypeString,
				Required: true,
			},
			"hec_endpoint_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"hec_token": {
				Type: schema.TypeString,
				Required: true,
			},
			"processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamProcessingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"retry_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamSplunkRetryOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"s3_backup_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"s3_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamS3DestinationConfiguration(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}
