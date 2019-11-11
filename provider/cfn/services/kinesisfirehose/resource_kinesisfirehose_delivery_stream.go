// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisFirehoseDeliveryStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisFirehoseDeliveryStreamCreate,
		Read:   resourceKinesisFirehoseDeliveryStreamRead,
		Update: resourceKinesisFirehoseDeliveryStreamUpdate,
		Delete: resourceKinesisFirehoseDeliveryStreamDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"delivery_stream_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"delivery_stream_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"elasticsearch_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamElasticsearchDestinationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"extended_s3_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamExtendedS3DestinationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_stream_source_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamKinesisStreamSourceConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"redshift_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamRedshiftDestinationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"s3_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamS3DestinationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"splunk_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamSplunkDestinationConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisFirehoseDeliveryStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisFirehose::DeliveryStream", ResourceKinesisFirehoseDeliveryStream(), data, meta)
}

func resourceKinesisFirehoseDeliveryStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisFirehose::DeliveryStream", ResourceKinesisFirehoseDeliveryStream(), data, meta)
}

func resourceKinesisFirehoseDeliveryStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisFirehose::DeliveryStream", ResourceKinesisFirehoseDeliveryStream(), data, meta)
}

func resourceKinesisFirehoseDeliveryStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisFirehose::DeliveryStream", data, meta)
}