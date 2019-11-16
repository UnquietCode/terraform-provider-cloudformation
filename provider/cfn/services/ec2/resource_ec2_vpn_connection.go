// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPNConnectionType string = "AWS::EC2::VPNConnection"

var eC2VPNConnectionProperties map[string]string = map[string]string{
	"customer_gateway_id": "CustomerGatewayId",
	"static_routes_only": "StaticRoutesOnly",
	"tags": "Tags",
	"transit_gateway_id": "TransitGatewayId",
	"type": "Type",
	"vpn_gateway_id": "VpnGatewayId",
	"vpn_tunnel_options_specifications": "VpnTunnelOptionsSpecifications",
}

func ResourceEC2VPNConnection() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPNConnectionExists,
		Read: resourceEC2VPNConnectionRead,
		Create: resourceEC2VPNConnectionCreate,
		Update: resourceEC2VPNConnectionUpdate,
		Delete: resourceEC2VPNConnectionDelete,
		CustomizeDiff: resourceEC2VPNConnectionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"customer_gateway_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"static_routes_only": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"transit_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpn_tunnel_options_specifications": {
				Type: schema.TypeSet,
				Elem: propertyVPNConnectionVpnTunnelOptionsSpecification(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPNConnectionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPNConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, eC2VPNConnectionProperties, meta)
}

func resourceEC2VPNConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, eC2VPNConnectionProperties, meta)
}

func resourceEC2VPNConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPNConnectionType, data, meta)
}

func resourceEC2VPNConnectionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPNConnectionType, data, meta)
}
