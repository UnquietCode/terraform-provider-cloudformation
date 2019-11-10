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

func ResourceSubnetCidrBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetCidrBlockCreate,
		Read:   resourceSubnetCidrBlockRead,
		Update: resourceSubnetCidrBlockUpdate,
		Delete: resourceSubnetCidrBlockDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSubnetCidrBlockCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetCidrBlock", data, meta)
}

func resourceSubnetCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetCidrBlock", data, meta)
}

func resourceSubnetCidrBlockUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetCidrBlock", data, meta)
}

func resourceSubnetCidrBlockDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetCidrBlock", data, meta)
}