// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDetectorModelAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iot_events": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelIotEvents(),
				Required: false,
				MaxItems: 1,
			},
			"reset_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelResetTimer(),
				Required: false,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSqs(),
				Required: false,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelFirehose(),
				Required: false,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSns(),
				Required: false,
				MaxItems: 1,
			},
			"iot_topic_publish": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelIotTopicPublish(),
				Required: false,
				MaxItems: 1,
			},
			"set_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSetTimer(),
				Required: false,
				MaxItems: 1,
			},
			"clear_timer": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelClearTimer(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelLambda(),
				Required: false,
				MaxItems: 1,
			},
			"set_variable": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelSetVariable(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}