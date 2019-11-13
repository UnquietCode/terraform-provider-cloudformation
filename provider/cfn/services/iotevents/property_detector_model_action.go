// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html

package iotevents

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDetectorModelAction(extras...string) *schema.Resource {
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
			"iot_events": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelIotEvents(),
				Optional: true,
				MaxItems: 1,
			},
			"reset_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelResetTimer(),
				Optional: true,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSqs(),
				Optional: true,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelFirehose(),
				Optional: true,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSns(),
				Optional: true,
				MaxItems: 1,
			},
			"iot_topic_publish": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelIotTopicPublish(),
				Optional: true,
				MaxItems: 1,
			},
			"set_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSetTimer(),
				Optional: true,
				MaxItems: 1,
			},
			"clear_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelClearTimer(),
				Optional: true,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelLambda(),
				Optional: true,
				MaxItems: 1,
			},
			"set_variable": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSetVariable(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}