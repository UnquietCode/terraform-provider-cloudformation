// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilter.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2TrafficMirrorFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2TrafficMirrorFilterCreate,
		Read:   resourceEC2TrafficMirrorFilterRead,
		Update: resourceEC2TrafficMirrorFilterUpdate,
		Delete: resourceEC2TrafficMirrorFilterDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network_services": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceEC2TrafficMirrorFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TrafficMirrorFilter", ResourceEC2TrafficMirrorFilter(), data, meta)
}

func resourceEC2TrafficMirrorFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TrafficMirrorFilter", ResourceEC2TrafficMirrorFilter(), data, meta)
}

func resourceEC2TrafficMirrorFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TrafficMirrorFilter", ResourceEC2TrafficMirrorFilter(), data, meta)
}

func resourceEC2TrafficMirrorFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TrafficMirrorFilter", data, meta)
}