// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const redshiftClusterParameterGroupType string = "AWS::Redshift::ClusterParameterGroup"

func ResourceRedshiftClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceRedshiftClusterParameterGroupRead,
		Create: resourceRedshiftClusterParameterGroupCreate,
		Update: resourceRedshiftClusterParameterGroupUpdate,
		Delete: resourceRedshiftClusterParameterGroupDelete,
		CustomizeDiff: resourceRedshiftClusterParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameter_group_family": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeList,
				Elem: propertyClusterParameterGroupParameter(),
				Optional: true,
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

func resourceRedshiftClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(redshiftClusterParameterGroupType, data, meta)
}

func resourceRedshiftClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(redshiftClusterParameterGroupType, data, meta)
}
