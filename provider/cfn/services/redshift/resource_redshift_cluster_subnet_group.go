// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const redshiftClusterSubnetGroupType string = "AWS::Redshift::ClusterSubnetGroup"

func ResourceRedshiftClusterSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRedshiftClusterSubnetGroupExists,
		Read: resourceRedshiftClusterSubnetGroupRead,
		Create: resourceRedshiftClusterSubnetGroupCreate,
		Update: resourceRedshiftClusterSubnetGroupUpdate,
		Delete: resourceRedshiftClusterSubnetGroupDelete,
		CustomizeDiff: resourceRedshiftClusterSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
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

func resourceRedshiftClusterSubnetGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRedshiftClusterSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterSubnetGroupType, ResourceRedshiftClusterSubnetGroup(), data, meta)
}

func resourceRedshiftClusterSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(redshiftClusterSubnetGroupType, ResourceRedshiftClusterSubnetGroup(), data, meta)
}

func resourceRedshiftClusterSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(redshiftClusterSubnetGroupType, ResourceRedshiftClusterSubnetGroup(), data, meta)
}

func resourceRedshiftClusterSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(redshiftClusterSubnetGroupType, data, meta)
}

func resourceRedshiftClusterSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(redshiftClusterSubnetGroupType, data, meta)
}
