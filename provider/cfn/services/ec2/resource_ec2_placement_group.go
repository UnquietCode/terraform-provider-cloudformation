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

func ResourceEC2PlacementGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2PlacementGroupCreate,
		Read:   resourceEC2PlacementGroupRead,
		Update: resourceEC2PlacementGroupUpdate,
		Delete: resourceEC2PlacementGroupDelete,

		Schema: map[string]*schema.Schema{
			"strategy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2PlacementGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::PlacementGroup", data, meta)
}

func resourceEC2PlacementGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::PlacementGroup", data, meta)
}

func resourceEC2PlacementGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::PlacementGroup", data, meta)
}

func resourceEC2PlacementGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::PlacementGroup", data, meta)
}