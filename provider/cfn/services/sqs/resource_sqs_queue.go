// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html

package sqs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sQSQueueType string = "AWS::SQS::Queue"

func ResourceSQSQueue() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"tags": misc.PropertyTags(),
			"visibility_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSQSQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sQSQueueType, ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sQSQueueType, ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sQSQueueType, ResourceSQSQueue(), data, meta)
}

func resourceSQSQueueDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sQSQueueType, data, meta)
}

func resourceSQSQueueCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sQSQueueType, data, meta)
}
