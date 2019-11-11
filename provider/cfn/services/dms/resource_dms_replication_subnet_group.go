// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSReplicationSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMSReplicationSubnetGroupCreate,
		Read:   resourceDMSReplicationSubnetGroupRead,
		Update: resourceDMSReplicationSubnetGroupUpdate,
		Delete: resourceDMSReplicationSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"replication_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_subnet_group_identifier": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSReplicationSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationSubnetGroup", ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationSubnetGroup", ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationSubnetGroup", ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationSubnetGroup", data, meta)
}