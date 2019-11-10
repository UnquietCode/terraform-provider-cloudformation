// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransitGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransitGatewayAttachmentCreate,
		Read:   resourceTransitGatewayAttachmentRead,
		Update: resourceTransitGatewayAttachmentUpdate,
		Delete: resourceTransitGatewayAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"transit_gateway_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceTransitGatewayAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::TransitGatewayAttachment", data, meta)
}

func resourceTransitGatewayAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::TransitGatewayAttachment", data, meta)
}

func resourceTransitGatewayAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::TransitGatewayAttachment", data, meta)
}

func resourceTransitGatewayAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::TransitGatewayAttachment", data, meta)
}