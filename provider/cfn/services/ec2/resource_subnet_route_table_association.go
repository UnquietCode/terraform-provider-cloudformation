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

func ResourceSubnetRouteTableAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetRouteTableAssociationCreate,
		Read:   resourceSubnetRouteTableAssociationRead,
		Update: resourceSubnetRouteTableAssociationUpdate,
		Delete: resourceSubnetRouteTableAssociationDelete,

		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSubnetRouteTableAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceSubnetRouteTableAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceSubnetRouteTableAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceSubnetRouteTableAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}