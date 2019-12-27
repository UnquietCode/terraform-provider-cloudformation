// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html

package kinesisfirehose

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisFirehoseDeliveryStreamType string = "AWS::KinesisFirehose::DeliveryStream"

func DatasourceKinesisFirehoseDeliveryStream() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKinesisFirehoseDeliveryStreamRead,
		
		Schema: map[string]*schema.Schema{
			"delivery_stream_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"delivery_stream_type": {
				Type: schema.TypeString,
				Optional: true,
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
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceKinesisFirehoseDeliveryStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisFirehoseDeliveryStreamType, DatasourceKinesisFirehoseDeliveryStream(), data, meta)
}
