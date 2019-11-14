// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html

package sqs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSQSQueue() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSQSQueueExists,
		Read: resourceSQSQueueRead,
		Create: resourceSQSQueueCreate,
		Update: resourceSQSQueueUpdate,
		Delete: resourceSQSQueueDelete,
		CustomizeDiff: resourceSQSQueueCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"content_based_deduplication": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"delay_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"fifo_queue": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"kms_data_key_reuse_period_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"kms_master_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"maximum_message_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"message_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"queue_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"receive_message_wait_time_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"redrive_policy": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"visibility_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSQSQueueExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSQSQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SQS::Queue", ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SQS::Queue", ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SQS::Queue", ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SQS::Queue", data, meta)
}

func resourceSQSQueueCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

