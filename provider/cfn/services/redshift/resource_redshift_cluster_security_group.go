// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceRedshiftClusterSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedshiftClusterSecurityGroupCreate,
		Read:   resourceRedshiftClusterSecurityGroupRead,
		Update: resourceRedshiftClusterSecurityGroupUpdate,
		Delete: resourceRedshiftClusterSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceRedshiftClusterSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::ClusterSecurityGroup", ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::ClusterSecurityGroup", ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::ClusterSecurityGroup", ResourceRedshiftClusterSecurityGroup(), data, meta)
}

func resourceRedshiftClusterSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::ClusterSecurityGroup", data, meta)
}