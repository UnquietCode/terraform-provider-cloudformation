// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRedshiftDestinationConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_watch_logging_options": {
				Type: schema.TypeList,
				Elem: propertyCloudWatchLoggingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"cluster_jdbcurl": {
				Type: schema.TypeString,
				Required: true,
			},
			"copy_command": {
				Type: schema.TypeList,
				Elem: propertyCopyCommand(),
				Required: true,
				MaxItems: 1,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
			},
			"processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyProcessingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_configuration": {
				Type: schema.TypeList,
				Elem: propertyS3DestinationConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"username": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}