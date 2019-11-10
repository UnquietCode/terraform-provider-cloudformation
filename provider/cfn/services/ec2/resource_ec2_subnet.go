// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2Subnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SubnetCreate,
		Read:   resourceEC2SubnetRead,
		Update: resourceEC2SubnetUpdate,
		Delete: resourceEC2SubnetDelete,

		Schema: map[string]*schema.Schema{
			"assign_ipv6_address_on_creation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipv6_cidr_block": {
				Type: schema.TypeString,
				Required: false,
			},
			"map_public_ip_on_launch": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2SubnetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Subnet", data, meta)
}

func resourceEC2SubnetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Subnet", data, meta)
}

func resourceEC2SubnetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Subnet", data, meta)
}

func resourceEC2SubnetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Subnet", data, meta)
}