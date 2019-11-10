// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sqs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceQueueCreate,
		Read:   resourceQueueRead,
		Update: resourceQueueUpdate,
		Delete: resourceQueueDelete,

		Schema: map[string]*schema.Schema{
			"content_based_deduplication": {
				Type: schema.TypeBool,
				Required: false,
			},
			"delay_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"fifo_queue": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"kms_data_key_reuse_period_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"kms_master_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"maximum_message_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"message_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"queue_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"receive_message_wait_time_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"redrive_policy": {
				Type: schema.TypeMap,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"visibility_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceQueueCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SQS::Queue", data, meta)
}

func resourceQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SQS::Queue", data, meta)
}

func resourceQueueUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SQS::Queue", data, meta)
}

func resourceQueueDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SQS::Queue", data, meta)
}