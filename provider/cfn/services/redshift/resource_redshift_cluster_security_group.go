// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const redshiftClusterSecurityGroupType string = "AWS::Redshift::ClusterSecurityGroup"

func ResourceRedshiftClusterSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRedshiftClusterSecurityGroupExists,
		Read: resourceRedshiftClusterSecurityGroupRead,
		Create: resourceRedshiftClusterSecurityGroupCreate,
		Update: resourceRedshiftClusterSecurityGroupUpdate,
		Delete: resourceRedshiftClusterSecurityGroupDelete,
		CustomizeDiff: resourceRedshiftClusterSecurityGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
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

func resourceRedshiftClusterSecurityGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRedshiftClusterSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterSecurityGroupType, ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(redshiftClusterSecurityGroupType, ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(redshiftClusterSecurityGroupType, ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(redshiftClusterSecurityGroupType, data, meta)
}

func resourceRedshiftClusterSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(redshiftClusterSecurityGroupType, data, meta)
}
