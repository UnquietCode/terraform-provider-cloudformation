// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2ClientVpnRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2ClientVpnRouteCreate,
		Read:   resourceEC2ClientVpnRouteRead,
		Delete: resourceEC2ClientVpnRouteDelete,

		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_vpc_subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2ClientVpnRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::ClientVpnRoute", ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::ClientVpnRoute", ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::ClientVpnRoute", ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::ClientVpnRoute", data, meta)
}