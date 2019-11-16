// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2CustomerGatewayType string = "AWS::EC2::CustomerGateway"

var eC2CustomerGatewayProperties map[string]string = map[string]string{
	"bgp_asn": "BgpAsn",
	"ip_address": "IpAddress",
	"tags": "Tags",
	"type": "Type",
}

func ResourceEC2CustomerGateway() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2CustomerGatewayExists,
		Read: resourceEC2CustomerGatewayRead,
		Create: resourceEC2CustomerGatewayCreate,
		Update: resourceEC2CustomerGatewayUpdate,
		Delete: resourceEC2CustomerGatewayDelete,
		CustomizeDiff: resourceEC2CustomerGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"bgp_asn": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ip_address": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
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

func resourceEC2CustomerGatewayExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2CustomerGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, meta)
}

func resourceEC2CustomerGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, eC2CustomerGatewayProperties, meta)
}

func resourceEC2CustomerGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, eC2CustomerGatewayProperties, meta)
}

func resourceEC2CustomerGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2CustomerGatewayType, data, meta)
}

func resourceEC2CustomerGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2CustomerGatewayType, data, meta)
}
