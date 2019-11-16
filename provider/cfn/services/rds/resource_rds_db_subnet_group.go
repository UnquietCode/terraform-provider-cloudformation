// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBSubnetGroupType string = "AWS::RDS::DBSubnetGroup"

var rDSDBSubnetGroupProperties map[string]string = map[string]string{
	"db_subnet_group_description": "DBSubnetGroupDescription",
	"db_subnet_group_name": "DBSubnetGroupName",
	"subnet_ids": "SubnetIds",
	"tags": "Tags",
}

func ResourceRDSDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSDBSubnetGroupExists,
		Read: resourceRDSDBSubnetGroupRead,
		Create: resourceRDSDBSubnetGroupCreate,
		Update: resourceRDSDBSubnetGroupUpdate,
		Delete: resourceRDSDBSubnetGroupDelete,
		CustomizeDiff: resourceRDSDBSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRDSDBSubnetGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBSubnetGroupType, ResourceRDSDBSubnetGroup(), data, meta)
}

func resourceRDSDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBSubnetGroupType, ResourceRDSDBSubnetGroup(), data, rDSDBSubnetGroupProperties, meta)
}

func resourceRDSDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBSubnetGroupType, ResourceRDSDBSubnetGroup(), data, rDSDBSubnetGroupProperties, meta)
}

func resourceRDSDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBSubnetGroupType, data, meta)
}

func resourceRDSDBSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBSubnetGroupType, data, meta)
}
