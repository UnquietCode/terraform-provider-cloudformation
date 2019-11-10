// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dax

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDAXSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDAXSubnetGroupCreate,
		Read:   resourceDAXSubnetGroupRead,
		Update: resourceDAXSubnetGroupUpdate,
		Delete: resourceDAXSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceDAXSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DAX::SubnetGroup", data, meta)
}

func resourceDAXSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DAX::SubnetGroup", data, meta)
}

func resourceDAXSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DAX::SubnetGroup", data, meta)
}

func resourceDAXSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DAX::SubnetGroup", data, meta)
}