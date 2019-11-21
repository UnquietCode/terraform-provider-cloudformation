// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2CustomerGatewayType string = "AWS::EC2::CustomerGateway"

func ResourceEC2CustomerGateway() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2CustomerGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, meta)
}

func resourceEC2CustomerGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, meta)
}

func resourceEC2CustomerGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2CustomerGatewayType, ResourceEC2CustomerGateway(), data, meta)
}

func resourceEC2CustomerGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2CustomerGatewayType, data, meta)
}

func resourceEC2CustomerGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2CustomerGatewayType, data, meta)
}
