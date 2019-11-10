// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBSubnetGroupCreate,
		Read:   resourceDBSubnetGroupRead,
		Update: resourceDBSubnetGroupUpdate,
		Delete: resourceDBSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBSubnetGroup", data, meta)
}

func resourceDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBSubnetGroup", data, meta)
}

func resourceDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBSubnetGroup", data, meta)
}

func resourceDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBSubnetGroup", data, meta)
}