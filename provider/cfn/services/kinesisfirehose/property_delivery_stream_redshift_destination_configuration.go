// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-redshiftdestinationconfiguration.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamRedshiftDestinationConfiguration(extras...string) *schema.Resource {
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
			"cluster_jdbcurl": {
				Type: schema.TypeString,
				Required: true,
			},
			"copy_command": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamCopyCommand(),
				Required: true,
				MaxItems: 1,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
			},
			"processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamProcessingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamS3DestinationConfiguration(),
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
