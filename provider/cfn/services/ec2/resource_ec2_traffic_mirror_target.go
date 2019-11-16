// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TrafficMirrorTargetType string = "AWS::EC2::TrafficMirrorTarget"

var eC2TrafficMirrorTargetProperties map[string]string = map[string]string{
	"network_load_balancer_arn": "NetworkLoadBalancerArn",
	"description": "Description",
	"network_interface_id": "NetworkInterfaceId",
	"tags": "Tags",
}

func ResourceEC2TrafficMirrorTarget() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TrafficMirrorTargetExists,
		Read: resourceEC2TrafficMirrorTargetRead,
		Create: resourceEC2TrafficMirrorTargetCreate,
		Update: resourceEC2TrafficMirrorTargetUpdate,
		Delete: resourceEC2TrafficMirrorTargetDelete,
		CustomizeDiff: resourceEC2TrafficMirrorTargetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"network_load_balancer_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceEC2TrafficMirrorTargetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TrafficMirrorTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TrafficMirrorTargetType, ResourceEC2TrafficMirrorTarget(), data, meta)
}

func resourceEC2TrafficMirrorTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TrafficMirrorTargetType, ResourceEC2TrafficMirrorTarget(), data, eC2TrafficMirrorTargetProperties, meta)
}

func resourceEC2TrafficMirrorTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TrafficMirrorTargetType, ResourceEC2TrafficMirrorTarget(), data, eC2TrafficMirrorTargetProperties, meta)
}

func resourceEC2TrafficMirrorTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TrafficMirrorTargetType, data, meta)
}

func resourceEC2TrafficMirrorTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TrafficMirrorTargetType, data, meta)
}
