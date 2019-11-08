// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTTopicRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTTopicRuleCreate,
		Read:   resourceIoTTopicRuleRead,
		Update: resourceIoTTopicRuleUpdate,
		Delete: resourceIoTTopicRuleDelete,

		Schema: map[string]*schema.Schema{
			"rule_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"topic_rule_payload": {
				Type: schema.TypeList,
				Elem: propertyTopicRulePayload(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceIoTTopicRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::TopicRule", data, meta)
}

func resourceIoTTopicRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::TopicRule", data, meta)
}

func resourceIoTTopicRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::TopicRule", data, meta)
}

func resourceIoTTopicRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::TopicRule", data, meta)
}