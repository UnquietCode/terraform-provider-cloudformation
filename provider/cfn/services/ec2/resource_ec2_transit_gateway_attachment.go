// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2TransitGatewayAttachmentType string = "AWS::EC2::TransitGatewayAttachment"

var eC2TransitGatewayAttachmentProperties map[string]string = map[string]string{
	"transit_gateway_id": "TransitGatewayId",
	"vpc_id": "VpcId",
	"subnet_ids": "SubnetIds",
	"tags": "Tags",
}

func ResourceEC2TransitGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2TransitGatewayAttachmentExists,
		Read: resourceEC2TransitGatewayAttachmentRead,
		Create: resourceEC2TransitGatewayAttachmentCreate,
		Update: resourceEC2TransitGatewayAttachmentUpdate,
		Delete: resourceEC2TransitGatewayAttachmentDelete,
		CustomizeDiff: resourceEC2TransitGatewayAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"transit_gateway_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2TransitGatewayAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2TransitGatewayAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2TransitGatewayAttachmentType, ResourceEC2TransitGatewayAttachment(), data, meta)
}

func resourceEC2TransitGatewayAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2TransitGatewayAttachmentType, ResourceEC2TransitGatewayAttachment(), data, eC2TransitGatewayAttachmentProperties, meta)
}

func resourceEC2TransitGatewayAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2TransitGatewayAttachmentType, ResourceEC2TransitGatewayAttachment(), data, eC2TransitGatewayAttachmentProperties, meta)
}

func resourceEC2TransitGatewayAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2TransitGatewayAttachmentType, data, meta)
}

func resourceEC2TransitGatewayAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2TransitGatewayAttachmentType, data, meta)
}
