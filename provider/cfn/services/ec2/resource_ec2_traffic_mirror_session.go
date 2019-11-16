// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorsession.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TrafficMirrorSessionType string = "AWS::EC2::TrafficMirrorSession"

var eC2TrafficMirrorSessionProperties map[string]string = map[string]string{
	"traffic_mirror_target_id": "TrafficMirrorTargetId",
	"description": "Description",
	"session_number": "SessionNumber",
	"virtual_network_id": "VirtualNetworkId",
	"packet_length": "PacketLength",
	"network_interface_id": "NetworkInterfaceId",
	"traffic_mirror_filter_id": "TrafficMirrorFilterId",
	"tags": "Tags",
}

func ResourceEC2TrafficMirrorSession() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TrafficMirrorSessionExists,
		Read: resourceEC2TrafficMirrorSessionRead,
		Create: resourceEC2TrafficMirrorSessionCreate,
		Update: resourceEC2TrafficMirrorSessionUpdate,
		Delete: resourceEC2TrafficMirrorSessionDelete,
		CustomizeDiff: resourceEC2TrafficMirrorSessionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"traffic_mirror_target_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"session_number": {
				Type: schema.TypeInt,
				Required: true,
			},
			"virtual_network_id": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"packet_length": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"traffic_mirror_filter_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TrafficMirrorSessionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TrafficMirrorSessionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TrafficMirrorSessionType, ResourceEC2TrafficMirrorSession(), data, meta)
}

func resourceEC2TrafficMirrorSessionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TrafficMirrorSessionType, ResourceEC2TrafficMirrorSession(), data, eC2TrafficMirrorSessionProperties, meta)
}

func resourceEC2TrafficMirrorSessionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TrafficMirrorSessionType, ResourceEC2TrafficMirrorSession(), data, eC2TrafficMirrorSessionProperties, meta)
}

func resourceEC2TrafficMirrorSessionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TrafficMirrorSessionType, data, meta)
}

func resourceEC2TrafficMirrorSessionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TrafficMirrorSessionType, data, meta)
}
