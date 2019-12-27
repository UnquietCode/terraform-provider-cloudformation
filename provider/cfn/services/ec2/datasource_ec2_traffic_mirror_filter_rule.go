// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TrafficMirrorFilterRuleType string = "AWS::EC2::TrafficMirrorFilterRule"

func DatasourceEC2TrafficMirrorFilterRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2TrafficMirrorFilterRuleRead,
		
		Schema: map[string]*schema.Schema{
			"destination_port_range": {
				Type: schema.TypeList,
				Elem: propertyTrafficMirrorFilterRuleTrafficMirrorPortRange(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_port_range": {
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2TrafficMirrorFilterRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TrafficMirrorFilterRuleType, DatasourceEC2TrafficMirrorFilterRule(), data, meta)
}
