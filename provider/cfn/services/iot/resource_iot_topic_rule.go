// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTTopicRuleType string = "AWS::IoT::TopicRule"

var ioTTopicRuleProperties map[string]string = map[string]string{
	"rule_name": "RuleName",
	"topic_rule_payload": "TopicRulePayload",
}

func ResourceIoTTopicRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTTopicRuleExists,
		Read: resourceIoTTopicRuleRead,
		Create: resourceIoTTopicRuleCreate,
		Update: resourceIoTTopicRuleUpdate,
		Delete: resourceIoTTopicRuleDelete,
		CustomizeDiff: resourceIoTTopicRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"rule_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"topic_rule_payload": {
				Type: schema.TypeSet,
				Elem: propertyTopicRuleTopicRulePayload(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTTopicRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTTopicRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTTopicRuleType, ResourceIoTTopicRule(), data, meta)
}

func resourceIoTTopicRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTTopicRuleType, ResourceIoTTopicRule(), data, ioTTopicRuleProperties, meta)
}

func resourceIoTTopicRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTTopicRuleType, ResourceIoTTopicRule(), data, ioTTopicRuleProperties, meta)
}

func resourceIoTTopicRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTTopicRuleType, data, meta)
}

func resourceIoTTopicRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTTopicRuleType, data, meta)
}
