// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html

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
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloudwatch_alarm": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleCloudwatchAlarmAction(),
				Optional: true,
				MaxItems: 1,
			},
			"cloudwatch_metric": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleCloudwatchMetricAction(),
				Optional: true,
				MaxItems: 1,
			},
			"dynamo_db": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleDynamoDBAction(),
				Optional: true,
				MaxItems: 1,
			},
			"dynamo_d_bv2": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleDynamoDBv2Action(),
				Optional: true,
				MaxItems: 1,
			},
			"elasticsearch": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleElasticsearchAction(),
				Optional: true,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleFirehoseAction(),
				Optional: true,
				MaxItems: 1,
			},
			"iot_analytics": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleIotAnalyticsAction(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleKinesisAction(),
				Optional: true,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleLambdaAction(),
				Optional: true,
				MaxItems: 1,
			},
			"republish": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleRepublishAction(),
				Optional: true,
				MaxItems: 1,
			},
			"s3": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleS3Action(),
				Optional: true,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleSnsAction(),
				Optional: true,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleSqsAction(),
				Optional: true,
				MaxItems: 1,
			},
			"step_functions": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleStepFunctionsAction(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}