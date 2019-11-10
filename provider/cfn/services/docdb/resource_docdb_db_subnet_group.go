// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html

package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDocDBDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocDBDBSubnetGroupCreate,
		Read:   resourceDocDBDBSubnetGroupRead,
		Update: resourceDocDBDBSubnetGroupUpdate,
		Delete: resourceDocDBDBSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
		},
	}
}

func resourceDocDBDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DocDB::DBSubnetGroup", ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DocDB::DBSubnetGroup", ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DocDB::DBSubnetGroup", ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DocDB::DBSubnetGroup", data, meta)
}