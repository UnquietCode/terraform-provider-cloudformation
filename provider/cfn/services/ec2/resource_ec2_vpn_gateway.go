// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPNGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPNGatewayExists,
		Read: resourceEC2VPNGatewayRead,
		Create: resourceEC2VPNGatewayCreate,
		Update: resourceEC2VPNGatewayUpdate,
		Delete: resourceEC2VPNGatewayDelete,
		CustomizeDiff: resourceEC2VPNGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"amazon_side_asn": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"type": {
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

func resourceEC2VPNGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPNGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPNGateway", ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPNGateway", ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPNGateway", ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPNGateway", data, meta)
}

func resourceEC2VPNGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
