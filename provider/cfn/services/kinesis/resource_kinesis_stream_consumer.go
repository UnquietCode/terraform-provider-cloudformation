// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisStreamConsumer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisStreamConsumerExists,
		Read: resourceKinesisStreamConsumerRead,
		Create: resourceKinesisStreamConsumerCreate,
		Update: resourceKinesisStreamConsumerUpdate,
		Delete: resourceKinesisStreamConsumerDelete,
		CustomizeDiff: resourceKinesisStreamConsumerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"consumer_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stream_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisStreamConsumerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisStreamConsumerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::StreamConsumer", ResourceKinesisStreamConsumer(), data, meta)
}

func resourceKinesisStreamConsumerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::StreamConsumer", ResourceKinesisStreamConsumer(), data, meta)
}

func resourceKinesisStreamConsumerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::StreamConsumer", ResourceKinesisStreamConsumer(), data, meta)
}

func resourceKinesisStreamConsumerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceKinesisStreamConsumerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
