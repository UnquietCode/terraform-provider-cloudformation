// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SubnetCidrBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SubnetCidrBlockCreate,
		Read:   resourceEC2SubnetCidrBlockRead,
		Delete: resourceEC2SubnetCidrBlockDelete,

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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SubnetCidrBlockCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetCidrBlock", ResourceEC2SubnetCidrBlock(), data, meta)
}

func resourceEC2SubnetCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetCidrBlock", ResourceEC2SubnetCidrBlock(), data, meta)
}

func resourceEC2SubnetCidrBlockUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetCidrBlock", ResourceEC2SubnetCidrBlock(), data, meta)
}

func resourceEC2SubnetCidrBlockDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetCidrBlock", data, meta)
}