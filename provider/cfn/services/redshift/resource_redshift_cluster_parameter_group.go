// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRedshiftClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedshiftClusterParameterGroupCreate,
		Read:   resourceRedshiftClusterParameterGroupRead,
		Update: resourceRedshiftClusterParameterGroupUpdate,
		Delete: resourceRedshiftClusterParameterGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameter_group_family": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type: schema.TypeList,
				Elem: propertyClusterParameterGroupParameter(),
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceRedshiftClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::ClusterParameterGroup", ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::ClusterParameterGroup", ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::ClusterParameterGroup", ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::ClusterParameterGroup", data, meta)
}