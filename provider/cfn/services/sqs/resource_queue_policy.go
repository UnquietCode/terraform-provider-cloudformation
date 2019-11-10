// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sqs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceQueuePolicyCreate,
		Read:   resourceQueuePolicyRead,
		Update: resourceQueuePolicyUpdate,
		Delete: resourceQueuePolicyDelete,

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
		},
	}
}

func resourceQueuePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SQS::QueuePolicy", data, meta)
}

func resourceQueuePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SQS::QueuePolicy", data, meta)
}

func resourceQueuePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SQS::QueuePolicy", data, meta)
}

func resourceQueuePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SQS::QueuePolicy", data, meta)
}