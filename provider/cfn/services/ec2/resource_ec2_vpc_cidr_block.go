// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCCidrBlockType string = "AWS::EC2::VPCCidrBlock"

var eC2VPCCidrBlockProperties map[string]string = map[string]string{
	"amazon_provided_ipv6_cidr_block": "AmazonProvidedIpv6CidrBlock",
	"cidr_block": "CidrBlock",
	"vpc_id": "VpcId",
}

func ResourceEC2VPCCidrBlock() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPCCidrBlockExists,
		Read: resourceEC2VPCCidrBlockRead,
		Create: resourceEC2VPCCidrBlockCreate,
		Update: resourceEC2VPCCidrBlockUpdate,
		Delete: resourceEC2VPCCidrBlockDelete,
		CustomizeDiff: resourceEC2VPCCidrBlockCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"amazon_provided_ipv6_cidr_block": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cidr_block": {
				Type: schema.TypeString,
				Optional: true,
			},
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

func resourceEC2VPCCidrBlockExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPCCidrBlockRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCCidrBlockType, ResourceEC2VPCCidrBlock(), data, meta)
}

func resourceEC2VPCCidrBlockCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCCidrBlockType, ResourceEC2VPCCidrBlock(), data, eC2VPCCidrBlockProperties, meta)
}

func resourceEC2VPCCidrBlockUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCCidrBlockType, ResourceEC2VPCCidrBlock(), data, eC2VPCCidrBlockProperties, meta)
}

func resourceEC2VPCCidrBlockDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCCidrBlockType, data, meta)
}

func resourceEC2VPCCidrBlockCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCCidrBlockType, data, meta)
}
