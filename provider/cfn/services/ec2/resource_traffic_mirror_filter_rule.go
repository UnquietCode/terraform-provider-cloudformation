// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTrafficMirrorFilterRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrafficMirrorFilterRuleCreate,
		Read:   resourceTrafficMirrorFilterRuleRead,
		Update: resourceTrafficMirrorFilterRuleUpdate,
		Delete: resourceTrafficMirrorFilterRuleDelete,

		Schema: map[string]*schema.Schema{
			"destination_port_range": {
				Type: schema.TypeList,
				Elem: propertyTrafficMirrorFilterRuleTrafficMirrorPortRange(),
				Required: false,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_port_range": {
				Type: schema.TypeList,
				Elem: propertyTrafficMirrorFilterRuleTrafficMirrorPortRange(),
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

func resourceTrafficMirrorFilterRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceTrafficMirrorFilterRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceTrafficMirrorFilterRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}

func resourceTrafficMirrorFilterRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TrafficMirrorFilterRule", data, meta)
}