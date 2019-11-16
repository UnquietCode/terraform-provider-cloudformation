// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SubnetRouteTableAssociationType string = "AWS::EC2::SubnetRouteTableAssociation"

func ResourceEC2SubnetRouteTableAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2SubnetRouteTableAssociationExists,
		Read: resourceEC2SubnetRouteTableAssociationRead,
		Create: resourceEC2SubnetRouteTableAssociationCreate,
		Update: resourceEC2SubnetRouteTableAssociationUpdate,
		Delete: resourceEC2SubnetRouteTableAssociationDelete,
		CustomizeDiff: resourceEC2SubnetRouteTableAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
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

func resourceEC2SubnetRouteTableAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2SubnetRouteTableAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SubnetRouteTableAssociationType, ResourceEC2SubnetRouteTableAssociation(), data, meta)
}

func resourceEC2SubnetRouteTableAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2SubnetRouteTableAssociationType, ResourceEC2SubnetRouteTableAssociation(), data, meta)
}

func resourceEC2SubnetRouteTableAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2SubnetRouteTableAssociationType, ResourceEC2SubnetRouteTableAssociation(), data, meta)
}

func resourceEC2SubnetRouteTableAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2SubnetRouteTableAssociationType, data, meta)
}

func resourceEC2SubnetRouteTableAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2SubnetRouteTableAssociationType, data, meta)
}
