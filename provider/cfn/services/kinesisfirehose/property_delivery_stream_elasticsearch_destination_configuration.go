// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamElasticsearchDestinationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buffering_hints": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamElasticsearchBufferingHints(),
				Required: true,
				MaxItems: 1,
			},
			"cloud_watch_logging_options": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamCloudWatchLoggingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"domain_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"index_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"index_rotation_period": {
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
				Elem: propertyDeliveryStreamElasticsearchRetryOptions(),
				Required: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_backup_mode": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamS3DestinationConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"type_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}