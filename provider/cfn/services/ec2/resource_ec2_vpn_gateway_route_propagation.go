// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPNGatewayRoutePropagation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPNGatewayRoutePropagationCreate,
		Read:   resourceEC2VPNGatewayRoutePropagationRead,
		Update: resourceEC2VPNGatewayRoutePropagationUpdate,
		Delete: resourceEC2VPNGatewayRoutePropagationDelete,

		Schema: map[string]*schema.Schema{
			"route_table_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEC2VPNGatewayRoutePropagationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceEC2VPNGatewayRoutePropagationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceEC2VPNGatewayRoutePropagationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}

func resourceEC2VPNGatewayRoutePropagationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNGatewayRoutePropagation", data, meta)
}