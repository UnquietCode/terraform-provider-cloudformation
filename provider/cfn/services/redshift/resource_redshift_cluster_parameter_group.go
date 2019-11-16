// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

var redshiftClusterParameterGroupProperties map[string]string = map[string]string{
	"description": "Description",
	"parameter_group_family": "ParameterGroupFamily",
	"parameters": "Parameters",
	"tags": "Tags",
}

func ResourceRedshiftClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRedshiftClusterParameterGroupExists,
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

func resourceRedshiftClusterParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRedshiftClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, meta)
}

func resourceRedshiftClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, redshiftClusterParameterGroupProperties, meta)
}

func resourceRedshiftClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(redshiftClusterParameterGroupType, ResourceRedshiftClusterParameterGroup(), data, redshiftClusterParameterGroupProperties, meta)
}

func resourceRedshiftClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(redshiftClusterParameterGroupType, data, meta)
}

func resourceRedshiftClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(redshiftClusterParameterGroupType, data, meta)
}
