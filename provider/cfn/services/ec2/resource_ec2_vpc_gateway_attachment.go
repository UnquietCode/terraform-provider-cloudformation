// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCGatewayAttachmentType string = "AWS::EC2::VPCGatewayAttachment"

var eC2VPCGatewayAttachmentProperties map[string]string = map[string]string{
	"internet_gateway_id": "InternetGatewayId",
	"vpc_id": "VpcId",
	"vpn_gateway_id": "VpnGatewayId",
}

func ResourceEC2VPCGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPCGatewayAttachmentExists,
		Read: resourceEC2VPCGatewayAttachmentRead,
		Create: resourceEC2VPCGatewayAttachmentCreate,
		Update: resourceEC2VPCGatewayAttachmentUpdate,
		Delete: resourceEC2VPCGatewayAttachmentDelete,
		CustomizeDiff: resourceEC2VPCGatewayAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"internet_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPCGatewayAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPCGatewayAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCGatewayAttachmentType, ResourceEC2VPCGatewayAttachment(), data, meta)
}

func resourceEC2VPCGatewayAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCGatewayAttachmentType, ResourceEC2VPCGatewayAttachment(), data, eC2VPCGatewayAttachmentProperties, meta)
}

func resourceEC2VPCGatewayAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCGatewayAttachmentType, ResourceEC2VPCGatewayAttachment(), data, eC2VPCGatewayAttachmentProperties, meta)
}

func resourceEC2VPCGatewayAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCGatewayAttachmentType, data, meta)
}

func resourceEC2VPCGatewayAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCGatewayAttachmentType, data, meta)
}
