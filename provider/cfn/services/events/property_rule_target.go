// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html

package events

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRuleTarget(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
	    if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
	        count = i
	    }
	}
	
	if count >= 5 {
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"batch_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleBatchParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"ecs_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleEcsParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"input": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_transformer": {
				Type: schema.TypeList,
				Elem: propertyRuleInputTransformer(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleKinesisParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"run_command_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleRunCommandParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"sqs_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleSqsParameters(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
