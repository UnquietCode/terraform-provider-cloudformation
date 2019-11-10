// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html

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
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"rule_name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceIoTTopicRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::TopicRule", ResourceIoTTopicRule(), data, meta)
}

func resourceIoTTopicRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::TopicRule", ResourceIoTTopicRule(), data, meta)
}

func resourceIoTTopicRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::TopicRule", ResourceIoTTopicRule(), data, meta)
}

func resourceIoTTopicRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::TopicRule", data, meta)
}