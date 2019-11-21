// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NatGatewayType string = "AWS::EC2::NatGateway"

func ResourceEC2NatGateway() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2NatGatewayRead,
		Create: resourceEC2NatGatewayCreate,
		Update: resourceEC2NatGatewayUpdate,
		Delete: resourceEC2NatGatewayDelete,
		CustomizeDiff: resourceEC2NatGatewayCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2NatGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NatGatewayType, ResourceEC2NatGateway(), data, meta)
}

func resourceEC2NatGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NatGatewayType, ResourceEC2NatGateway(), data, meta)
}

func resourceEC2NatGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NatGatewayType, ResourceEC2NatGateway(), data, meta)
}

func resourceEC2NatGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NatGatewayType, data, meta)
}

func resourceEC2NatGatewayCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NatGatewayType, data, meta)
}
