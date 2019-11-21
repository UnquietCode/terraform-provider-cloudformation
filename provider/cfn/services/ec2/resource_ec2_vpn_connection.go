// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPNConnectionType string = "AWS::EC2::VPNConnection"

func ResourceEC2VPNConnection() *schema.Resource {
	return &schema.Resource{
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
			"tags": misc.PropertyTags(),
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2VPNConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPNConnectionType, ResourceEC2VPNConnection(), data, meta)
}

func resourceEC2VPNConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPNConnectionType, data, meta)
}

func resourceEC2VPNConnectionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPNConnectionType, data, meta)
}
