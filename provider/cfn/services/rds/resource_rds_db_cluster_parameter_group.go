// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBClusterParameterGroupType string = "AWS::RDS::DBClusterParameterGroup"

var rDSDBClusterParameterGroupProperties map[string]string = map[string]string{
	"description": "Description",
	"family": "Family",
	"parameters": "Parameters",
	"tags": "Tags",
}

func ResourceRDSDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSDBClusterParameterGroupExists,
		Read: resourceRDSDBClusterParameterGroupRead,
		Create: resourceRDSDBClusterParameterGroupCreate,
		Update: resourceRDSDBClusterParameterGroupUpdate,
		Delete: resourceRDSDBClusterParameterGroupDelete,
		CustomizeDiff: resourceRDSDBClusterParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
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

func resourceRDSDBClusterParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBClusterParameterGroupType, ResourceRDSDBClusterParameterGroup(), data, meta)
}

func resourceRDSDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBClusterParameterGroupType, ResourceRDSDBClusterParameterGroup(), data, rDSDBClusterParameterGroupProperties, meta)
}

func resourceRDSDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBClusterParameterGroupType, ResourceRDSDBClusterParameterGroup(), data, rDSDBClusterParameterGroupProperties, meta)
}

func resourceRDSDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBClusterParameterGroupType, data, meta)
}

func resourceRDSDBClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBClusterParameterGroupType, data, meta)
}
