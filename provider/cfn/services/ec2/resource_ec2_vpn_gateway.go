// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPNGatewayType string = "AWS::EC2::VPNGateway"

func ResourceEC2VPNGateway() *schema.Resource {
	return &schema.Resource{
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
			"tags": misc.PropertyTags(),
			"type": {
				Type: schema.TypeString,
				Required: true,
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

func resourceEC2VPNGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNGatewayType, ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPNGatewayType, ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPNGatewayType, ResourceEC2VPNGateway(), data, meta)
}

func resourceEC2VPNGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPNGatewayType, data, meta)
}

func resourceEC2VPNGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPNGatewayType, data, meta)
}
