// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTopicRuleAction(extras...string) *schema.Resource {
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
			"cloudwatch_alarm": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleCloudwatchAlarmAction(),
				Required: false,
				MaxItems: 1,
			},
			"cloudwatch_metric": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleCloudwatchMetricAction(),
				Required: false,
				MaxItems: 1,
			},
			"dynamo_db": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleDynamoDBAction(),
				Required: false,
				MaxItems: 1,
			},
			"dynamo_d_bv2": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleDynamoDBv2Action(),
				Required: false,
				MaxItems: 1,
			},
			"elasticsearch": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleElasticsearchAction(),
				Required: false,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleFirehoseAction(),
				Required: false,
				MaxItems: 1,
			},
			"iot_analytics": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleIotAnalyticsAction(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleKinesisAction(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleLambdaAction(),
				Required: false,
				MaxItems: 1,
			},
			"republish": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleRepublishAction(),
				Required: false,
				MaxItems: 1,
			},
			"s3": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleS3Action(),
				Required: false,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleSnsAction(),
				Required: false,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleSqsAction(),
				Required: false,
				MaxItems: 1,
			},
			"step_functions": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleStepFunctionsAction(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}