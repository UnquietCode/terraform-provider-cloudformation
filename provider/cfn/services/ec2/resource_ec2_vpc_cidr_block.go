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

func ResourceEC2VPCCidrBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPCCidrBlockCreate,
		Read:   resourceEC2VPCCidrBlockRead,
		Update: resourceEC2VPCCidrBlockUpdate,
		Delete: resourceEC2VPCCidrBlockDelete,

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

func resourceEC2VPCCidrBlockCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceEC2VPCCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceEC2VPCCidrBlockUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCCidrBlock", data, meta)
}

func resourceEC2VPCCidrBlockDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCCidrBlock", data, meta)
}