// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2ClientVpnRouteType string = "AWS::EC2::ClientVpnRoute"

func ResourceEC2ClientVpnRoute() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2ClientVpnRouteRead,
		Create: resourceEC2ClientVpnRouteCreate,
		Update: resourceEC2ClientVpnRouteUpdate,
		Delete: resourceEC2ClientVpnRouteDelete,
		CustomizeDiff: resourceEC2ClientVpnRouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_vpn_endpoint_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_vpc_subnet_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destination_cidr_block": {
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

func resourceEC2ClientVpnRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2ClientVpnRouteType, ResourceEC2ClientVpnRoute(), data, meta)
}

func resourceEC2ClientVpnRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2ClientVpnRouteType, data, meta)
}

func resourceEC2ClientVpnRouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2ClientVpnRouteType, data, meta)
}
