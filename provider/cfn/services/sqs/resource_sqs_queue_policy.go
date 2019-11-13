// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html

package sqs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSQSQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSQSQueuePolicyExists,
		Read:   resourceSQSQueuePolicyRead,
		Create: resourceSQSQueuePolicyCreate,
		Update: resourceSQSQueuePolicyUpdate,
		Delete: resourceSQSQueuePolicyDelete,
		CustomizeDiff: resourceSQSQueuePolicyCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"queues": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceSQSQueuePolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSQSQueuePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SQS::QueuePolicy", ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SQS::QueuePolicy", ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SQS::QueuePolicy", ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SQS::QueuePolicy", data, meta)
}

func resourceSQSQueuePolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}