// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TrafficMirrorFilterRuleType string = "AWS::EC2::TrafficMirrorFilterRule"

var eC2TrafficMirrorFilterRuleProperties map[string]string = map[string]string{
	"destination_port_range": "DestinationPortRange",
	"description": "Description",
	"source_port_range": "SourcePortRange",
	"rule_action": "RuleAction",
	"source_cidr_block": "SourceCidrBlock",
	"rule_number": "RuleNumber",
	"destination_cidr_block": "DestinationCidrBlock",
	"traffic_mirror_filter_id": "TrafficMirrorFilterId",
	"traffic_direction": "TrafficDirection",
	"protocol": "Protocol",
}

func ResourceEC2TrafficMirrorFilterRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TrafficMirrorFilterRuleExists,
		Read: resourceEC2TrafficMirrorFilterRuleRead,
		Create: resourceEC2TrafficMirrorFilterRuleCreate,
		Update: resourceEC2TrafficMirrorFilterRuleUpdate,
		Delete: resourceEC2TrafficMirrorFilterRuleDelete,
		CustomizeDiff: resourceEC2TrafficMirrorFilterRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"destination_port_range": {
				Type: schema.TypeSet,
				Elem: propertyTrafficMirrorFilterRuleTrafficMirrorPortRange(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_port_range": {
				Type: schema.TypeSet,
				Elem: propertyTrafficMirrorFilterRuleTrafficMirrorPortRange(),
				Optional: true,
				MaxItems: 1,
			},
			"rule_action": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"rule_number": {
				Type: schema.TypeInt,
				Required: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"traffic_mirror_filter_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"traffic_direction": {
				Type: schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TrafficMirrorFilterRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TrafficMirrorFilterRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TrafficMirrorFilterRuleType, ResourceEC2TrafficMirrorFilterRule(), data, meta)
}

func resourceEC2TrafficMirrorFilterRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TrafficMirrorFilterRuleType, ResourceEC2TrafficMirrorFilterRule(), data, eC2TrafficMirrorFilterRuleProperties, meta)
}

func resourceEC2TrafficMirrorFilterRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TrafficMirrorFilterRuleType, ResourceEC2TrafficMirrorFilterRule(), data, eC2TrafficMirrorFilterRuleProperties, meta)
}

func resourceEC2TrafficMirrorFilterRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TrafficMirrorFilterRuleType, data, meta)
}

func resourceEC2TrafficMirrorFilterRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TrafficMirrorFilterRuleType, data, meta)
}
