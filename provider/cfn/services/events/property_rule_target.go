// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"ecs_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleEcsParameters(),
				Required: false,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"input": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_transformer": {
				Type: schema.TypeList,
				Elem: propertyRuleInputTransformer(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleKinesisParameters(),
				Required: false,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"run_command_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleRunCommandParameters(),
				Required: false,
				MaxItems: 1,
			},
			"sqs_parameters": {
				Type: schema.TypeList,
				Elem: propertyRuleSqsParameters(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}