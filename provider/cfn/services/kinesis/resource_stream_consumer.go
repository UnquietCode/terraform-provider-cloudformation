// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStreamConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamConsumerCreate,
		Read:   resourceStreamConsumerRead,
		Update: resourceStreamConsumerUpdate,
		Delete: resourceStreamConsumerDelete,

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

func resourceStreamConsumerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceStreamConsumerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceStreamConsumerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::StreamConsumer", data, meta)
}

func resourceStreamConsumerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::StreamConsumer", data, meta)
}