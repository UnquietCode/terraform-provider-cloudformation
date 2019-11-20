// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const eC2SubnetType string = "AWS::EC2::Subnet"

func ResourceEC2Subnet() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2SubnetRead,
		Create: resourceEC2SubnetCreate,
		Update: resourceEC2SubnetUpdate,
		Delete: resourceEC2SubnetDelete,
		CustomizeDiff: resourceEC2SubnetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"assign_ipv6_address_on_creation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"ipv6_cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
			"map_public_ip_on_launch": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"vpc_id": {
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

func resourceEC2SubnetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SubnetType, ResourceEC2Subnet(), data, meta)
}

func resourceEC2SubnetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2SubnetType, ResourceEC2Subnet(), data, meta)
}

func resourceEC2SubnetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2SubnetType, ResourceEC2Subnet(), data, meta)
}

func resourceEC2SubnetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2SubnetType, data, meta)
}

func resourceEC2SubnetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2SubnetType, data, meta)
}
