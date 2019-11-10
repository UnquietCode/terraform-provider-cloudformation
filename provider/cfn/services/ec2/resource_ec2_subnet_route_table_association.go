// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2SubnetRouteTableAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2SubnetRouteTableAssociationCreate,
		Read:   resourceEC2SubnetRouteTableAssociationRead,
		Update: resourceEC2SubnetRouteTableAssociationUpdate,
		Delete: resourceEC2SubnetRouteTableAssociationDelete,

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

func resourceEC2SubnetRouteTableAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceEC2SubnetRouteTableAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceEC2SubnetRouteTableAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}

func resourceEC2SubnetRouteTableAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::SubnetRouteTableAssociation", data, meta)
}