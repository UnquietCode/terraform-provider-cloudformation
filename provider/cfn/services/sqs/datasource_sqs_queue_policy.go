// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html

package sqs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sQSQueuePolicyType string = "AWS::SQS::QueuePolicy"

func DatasourceSQSQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSQSQueuePolicyRead,
		
		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceSQSQueuePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sQSQueuePolicyType, DatasourceSQSQueuePolicy(), data, meta)
}
