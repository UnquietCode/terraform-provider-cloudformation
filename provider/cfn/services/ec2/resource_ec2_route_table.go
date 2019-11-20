// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2RouteTableType string = "AWS::EC2::RouteTable"

func ResourceEC2RouteTable() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2RouteTableRead,
		Create: resourceEC2RouteTableCreate,
		Update: resourceEC2RouteTableUpdate,
		Delete: resourceEC2RouteTableDelete,
		CustomizeDiff: resourceEC2RouteTableCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": misc.PropertyTags(),
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

func resourceEC2RouteTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2RouteTableType, ResourceEC2RouteTable(), data, meta)
}

func resourceEC2RouteTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2RouteTableType, ResourceEC2RouteTable(), data, meta)
}

func resourceEC2RouteTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2RouteTableType, ResourceEC2RouteTable(), data, meta)
}

func resourceEC2RouteTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2RouteTableType, data, meta)
}

func resourceEC2RouteTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2RouteTableType, data, meta)
}
