// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceEC2VPNConnection() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2VPNConnectionRead,
		
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEC2VPNConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNConnectionType, DatasourceEC2VPNConnection(), data, meta)
}
