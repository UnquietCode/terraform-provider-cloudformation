// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html

package dms

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dMSReplicationSubnetGroupType string = "AWS::DMS::ReplicationSubnetGroup"

func ResourceDMSReplicationSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceDMSReplicationSubnetGroupRead,
		Create: resourceDMSReplicationSubnetGroupCreate,
		Update: resourceDMSReplicationSubnetGroupUpdate,
		Delete: resourceDMSReplicationSubnetGroupDelete,
		CustomizeDiff: resourceDMSReplicationSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"replication_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_subnet_group_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceDMSReplicationSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSReplicationSubnetGroupType, ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dMSReplicationSubnetGroupType, ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dMSReplicationSubnetGroupType, ResourceDMSReplicationSubnetGroup(), data, meta)
}

func resourceDMSReplicationSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dMSReplicationSubnetGroupType, data, meta)
}

func resourceDMSReplicationSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dMSReplicationSubnetGroupType, data, meta)
}
