// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloudwatch_alarm": {
				Type: schema.TypeList,
				Elem: propertyCloudwatchAlarmAction(),
				Required: false,
				MaxItems: 1,
			},
			"cloudwatch_metric": {
				Type: schema.TypeList,
				Elem: propertyCloudwatchMetricAction(),
				Required: false,
				MaxItems: 1,
			},
			"dynamo_db": {
				Type: schema.TypeList,
				Elem: propertyDynamoDBAction(),
				Required: false,
				MaxItems: 1,
			},
			"dynamo_d_bv2": {
				Type: schema.TypeList,
				Elem: propertyDynamoDBv2Action(),
				Required: false,
				MaxItems: 1,
			},
			"elasticsearch": {
				Type: schema.TypeList,
				Elem: propertyElasticsearchAction(),
				Required: false,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyFirehoseAction(),
				Required: false,
				MaxItems: 1,
			},
			"iot_analytics": {
				Type: schema.TypeList,
				Elem: propertyIotAnalyticsAction(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis": {
				Type: schema.TypeList,
				Elem: propertyKinesisAction(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyLambdaAction(),
				Required: false,
				MaxItems: 1,
			},
			"republish": {
				Type: schema.TypeList,
				Elem: propertyRepublishAction(),
				Required: false,
				MaxItems: 1,
			},
			"s3": {
				Type: schema.TypeList,
				Elem: propertyS3Action(),
				Required: false,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertySnsAction(),
				Required: false,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertySqsAction(),
				Required: false,
				MaxItems: 1,
			},
			"step_functions": {
				Type: schema.TypeList,
				Elem: propertyStepFunctionsAction(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}