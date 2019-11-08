// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisStreamConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisStreamConsumerCreate,
		Read:   resourceKinesisStreamConsumerRead,
		Update: resourceKinesisStreamConsumerUpdate,
		Delete: resourceKinesisStreamConsumerDelete,

		Schema: map[string]*schema.Schema{
			"consumer_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stream_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisStreamConsumerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceKinesisStreamConsumerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceKinesisStreamConsumerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceKinesisStreamConsumerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::StreamConsumer", data, meta)
}