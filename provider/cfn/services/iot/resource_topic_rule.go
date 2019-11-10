// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTopicRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTopicRuleCreate,
		Read:   resourceTopicRuleRead,
		Update: resourceTopicRuleUpdate,
		Delete: resourceTopicRuleDelete,

		Schema: map[string]*schema.Schema{
			"rule_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"topic_rule_payload": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleTopicRulePayload(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceTopicRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::TopicRule", data, meta)
}

func resourceTopicRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::TopicRule", data, meta)
}

func resourceTopicRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::TopicRule", data, meta)
}

func resourceTopicRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::TopicRule", data, meta)
}