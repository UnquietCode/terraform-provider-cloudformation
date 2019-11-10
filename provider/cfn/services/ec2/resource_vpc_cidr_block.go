// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVPCCidrBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCCidrBlockCreate,
		Read:   resourceVPCCidrBlockRead,
		Update: resourceVPCCidrBlockUpdate,
		Delete: resourceVPCCidrBlockDelete,

		Schema: map[string]*schema.Schema{
			"amazon_provided_ipv6_cidr_block": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"cidr_block": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCCidrBlockCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceVPCCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceVPCCidrBlockUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceVPCCidrBlockDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCCidrBlock", data, meta)
}