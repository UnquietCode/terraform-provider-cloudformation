// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPCGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPCGatewayAttachmentCreate,
		Read:   resourceEC2VPCGatewayAttachmentRead,
		Update: resourceEC2VPCGatewayAttachmentUpdate,
		Delete: resourceEC2VPCGatewayAttachmentDelete,

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
		},
	}
}

func resourceEC2VPCGatewayAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceEC2VPCGatewayAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceEC2VPCGatewayAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceEC2VPCGatewayAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCGatewayAttachment", data, meta)
}