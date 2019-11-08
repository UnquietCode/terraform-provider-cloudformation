// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TrafficMirrorFilterRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2TrafficMirrorFilterRuleCreate,
		Read:   resourceEC2TrafficMirrorFilterRuleRead,
		Update: resourceEC2TrafficMirrorFilterRuleUpdate,
		Delete: resourceEC2TrafficMirrorFilterRuleDelete,

		Schema: map[string]*schema.Schema{
			"destination_port_range": {
				Type: schema.TypeList,
				Elem: propertyTrafficMirrorPortRange(),
				Required: false,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_port_range": {
				Type: schema.TypeList,
				Elem: propertyTrafficMirrorPortRange(),
				Required: false,
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
				ForceNew: true,
			},
			"traffic_direction": {
				Type: schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceEC2TrafficMirrorFilterRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceEC2TrafficMirrorFilterRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceEC2TrafficMirrorFilterRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceEC2TrafficMirrorFilterRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}