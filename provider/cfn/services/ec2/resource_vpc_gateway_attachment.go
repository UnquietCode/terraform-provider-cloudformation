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

func ResourceVPCGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCGatewayAttachmentCreate,
		Read:   resourceVPCGatewayAttachmentRead,
		Update: resourceVPCGatewayAttachmentUpdate,
		Delete: resourceVPCGatewayAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"internet_gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpn_gateway_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceVPCGatewayAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceVPCGatewayAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceVPCGatewayAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCGatewayAttachment", data, meta)
}

func resourceVPCGatewayAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCGatewayAttachment", data, meta)
}