// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceEC2VPNConnection() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPNConnectionExists,
		Read:   resourceEC2VPNConnectionRead,
		Create: resourceEC2VPNConnectionCreate,
		Update: resourceEC2VPNConnectionUpdate,
		Delete: resourceEC2VPNConnectionDelete,
		
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
	return plugin.ResourceRead("AWS::EC2::VPNConnection", ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNConnection", ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNConnection", ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNConnection", data, meta)
}