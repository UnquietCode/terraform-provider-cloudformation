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

func ResourcePlacementGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlacementGroupCreate,
		Read:   resourcePlacementGroupRead,
		Update: resourcePlacementGroupUpdate,
		Delete: resourcePlacementGroupDelete,

		Schema: map[string]*schema.Schema{
			"strategy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourcePlacementGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::PlacementGroup", data, meta)
}

func resourcePlacementGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::PlacementGroup", data, meta)
}

func resourcePlacementGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::PlacementGroup", data, meta)
}

func resourcePlacementGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::PlacementGroup", data, meta)
}