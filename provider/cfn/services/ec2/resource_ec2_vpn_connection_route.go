// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPNConnectionRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPNConnectionRouteCreate,
		Read:   resourceEC2VPNConnectionRouteRead,
		Delete: resourceEC2VPNConnectionRouteDelete,

		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpn_connection_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPNConnectionRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNConnectionRoute", ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNConnectionRoute", ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNConnectionRoute", ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNConnectionRoute", data, meta)
}