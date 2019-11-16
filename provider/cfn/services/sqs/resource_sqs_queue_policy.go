// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html

package sqs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sQSQueuePolicyType string = "AWS::SQS::QueuePolicy"

func ResourceSQSQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSQSQueuePolicyExists,
		Read: resourceSQSQueuePolicyRead,
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
	return plugin.ResourceRead(sQSQueuePolicyType, ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sQSQueuePolicyType, ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sQSQueuePolicyType, ResourceSQSQueuePolicy(), data, meta)
}

func resourceSQSQueuePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sQSQueuePolicyType, data, meta)
}

func resourceSQSQueuePolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sQSQueuePolicyType, data, meta)
}
