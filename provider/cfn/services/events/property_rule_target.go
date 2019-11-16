// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html

package events

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var ruleTargetProperties map[string]string = map[string]string{
	"arn": "Arn",
	"ecs_parameters": "EcsParameters",
	"id": "Id",
	"input": "Input",
	"input_path": "InputPath",
	"input_transformer": "InputTransformer",
	"kinesis_parameters": "KinesisParameters",
	"role_arn": "RoleArn",
	"run_command_parameters": "RunCommandParameters",
	"sqs_parameters": "SqsParameters",
}

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
			"ecs_parameters": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyRuleInputTransformer(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_parameters": {
				Type: schema.TypeSet,
				Elem: propertyRuleKinesisParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"run_command_parameters": {
				Type: schema.TypeSet,
				Elem: propertyRuleRunCommandParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"sqs_parameters": {
				Type: schema.TypeSet,
				Elem: propertyRuleSqsParameters(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
