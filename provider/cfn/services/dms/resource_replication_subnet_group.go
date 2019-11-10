// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceReplicationSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceReplicationSubnetGroupCreate,
		Read:   resourceReplicationSubnetGroupRead,
		Update: resourceReplicationSubnetGroupUpdate,
		Delete: resourceReplicationSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"replication_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_subnet_group_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
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
				ForceNew: true,
			},
		},
	}
}

func resourceReplicationSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationSubnetGroup", data, meta)
}

func resourceReplicationSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationSubnetGroup", data, meta)
}

func resourceReplicationSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationSubnetGroup", data, meta)
}

func resourceReplicationSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationSubnetGroup", data, meta)
}