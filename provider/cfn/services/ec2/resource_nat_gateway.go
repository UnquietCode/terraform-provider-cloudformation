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

func ResourceNatGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceNatGatewayCreate,
		Read:   resourceNatGatewayRead,
		Update: resourceNatGatewayUpdate,
		Delete: resourceNatGatewayDelete,

		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceNatGatewayCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NatGateway", data, meta)
}

func resourceNatGatewayRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NatGateway", data, meta)
}

func resourceNatGatewayUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NatGateway", data, meta)
}

func resourceNatGatewayDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NatGateway", data, meta)
}