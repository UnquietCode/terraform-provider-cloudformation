// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TrafficMirrorTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2TrafficMirrorTargetCreate,
		Read:   resourceEC2TrafficMirrorTargetRead,
		Update: resourceEC2TrafficMirrorTargetUpdate,
		Delete: resourceEC2TrafficMirrorTargetDelete,

		Schema: map[string]*schema.Schema{
			"network_load_balancer_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceEC2TrafficMirrorTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TrafficMirrorTarget", data, meta)
}

func resourceEC2TrafficMirrorTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TrafficMirrorTarget", data, meta)
}

func resourceEC2TrafficMirrorTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TrafficMirrorTarget", data, meta)
}

func resourceEC2TrafficMirrorTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TrafficMirrorTarget", data, meta)
}