// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const docDBDBSubnetGroupType string = "AWS::DocDB::DBSubnetGroup"

func ResourceDocDBDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceDocDBDBSubnetGroupRead,
		Create: resourceDocDBDBSubnetGroupCreate,
		Update: resourceDocDBDBSubnetGroupUpdate,
		Delete: resourceDocDBDBSubnetGroupDelete,
		CustomizeDiff: resourceDocDBDBSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
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
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDocDBDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBSubnetGroupType, ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBSubnetGroupType, ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBSubnetGroupType, ResourceDocDBDBSubnetGroup(), data, meta)
}

func resourceDocDBDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBSubnetGroupType, data, meta)
}

func resourceDocDBDBSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBSubnetGroupType, data, meta)
}
