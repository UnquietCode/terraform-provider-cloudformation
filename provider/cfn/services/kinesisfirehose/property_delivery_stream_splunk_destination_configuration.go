// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamSplunkDestinationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_watch_logging_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamCloudWatchLoggingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"hec_acknowledgment_timeout_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
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
				Required: false,
				MaxItems: 1,
			},
			"retry_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamSplunkRetryOptions(),
				Required: false,
				MaxItems: 1,
			},
			"s3_backup_mode": {
				Type: schema.TypeString,
				Required: false,
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