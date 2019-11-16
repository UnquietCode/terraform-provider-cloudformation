// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnRouteType string = "AWS::EC2::ClientVpnRoute"

var eC2ClientVpnRouteProperties map[string]string = map[string]string{
	"client_vpn_endpoint_id": "ClientVpnEndpointId",
	"target_vpc_subnet_id": "TargetVpcSubnetId",
	"description": "Description",
	"destination_cidr_block": "DestinationCidrBlock",
}

func ResourceEC2ClientVpnRoute() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2ClientVpnRouteExists,
		Read: resourceEC2ClientVpnRouteRead,
		Create: resourceEC2ClientVpnRouteCreate,
		Update: resourceEC2ClientVpnRouteUpdate,
		Delete: resourceEC2ClientVpnRouteDelete,
		CustomizeDiff: resourceEC2ClientVpnRouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_vpc_subnet_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2ClientVpnRouteExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2ClientVpnRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, eC2ClientVpnRouteProperties, meta)
}

func resourceEC2ClientVpnRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, eC2ClientVpnRouteProperties, meta)
}

func resourceEC2ClientVpnRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2ClientVpnRouteType, data, meta)
}

func resourceEC2ClientVpnRouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2ClientVpnRouteType, data, meta)
}
