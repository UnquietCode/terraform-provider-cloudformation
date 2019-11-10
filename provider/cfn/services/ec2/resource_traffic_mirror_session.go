// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTrafficMirrorSession() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrafficMirrorSessionCreate,
		Read:   resourceTrafficMirrorSessionRead,
		Update: resourceTrafficMirrorSessionUpdate,
		Delete: resourceTrafficMirrorSessionDelete,

		Schema: map[string]*schema.Schema{
			"traffic_mirror_target_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"session_number": {
				Type: schema.TypeInt,
				Required: true,
			},
			"virtual_network_id": {
				Type: schema.TypeInt,
				Required: false,
			},
			"packet_length": {
				Type: schema.TypeInt,
				Required: false,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"traffic_mirror_filter_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceTrafficMirrorSessionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TrafficMirrorSession", data, meta)
}

func resourceTrafficMirrorSessionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TrafficMirrorSession", data, meta)
}

func resourceTrafficMirrorSessionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TrafficMirrorSession", data, meta)
}

func resourceTrafficMirrorSessionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TrafficMirrorSession", data, meta)
}