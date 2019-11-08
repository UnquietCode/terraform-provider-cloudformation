// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iot_events": {
				Type: schema.TypeList,
				Elem: propertyIotEvents(),
				Required: false,
				MaxItems: 1,
			},
			"reset_timer": {
				Type: schema.TypeList,
				Elem: propertyResetTimer(),
				Required: false,
				MaxItems: 1,
			},
			"sqs": {
				Type: schema.TypeList,
				Elem: propertySqs(),
				Required: false,
				MaxItems: 1,
			},
			"firehose": {
				Type: schema.TypeList,
				Elem: propertyFirehose(),
				Required: false,
				MaxItems: 1,
			},
			"sns": {
				Type: schema.TypeList,
				Elem: propertySns(),
				Required: false,
				MaxItems: 1,
			},
			"iot_topic_publish": {
				Type: schema.TypeList,
				Elem: propertyIotTopicPublish(),
				Required: false,
				MaxItems: 1,
			},
			"set_timer": {
				Type: schema.TypeList,
				Elem: propertySetTimer(),
				Required: false,
				MaxItems: 1,
			},
			"clear_timer": {
				Type: schema.TypeList,
				Elem: propertyClearTimer(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyLambda(),
				Required: false,
				MaxItems: 1,
			},
			"set_variable": {
				Type: schema.TypeList,
				Elem: propertySetVariable(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}