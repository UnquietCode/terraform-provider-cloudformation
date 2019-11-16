// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html

package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const docDBDBClusterParameterGroupType string = "AWS::DocDB::DBClusterParameterGroup"

func ResourceDocDBDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDocDBDBClusterParameterGroupExists,
		Read: resourceDocDBDBClusterParameterGroupRead,
		Create: resourceDocDBDBClusterParameterGroupCreate,
		Update: resourceDocDBDBClusterParameterGroupUpdate,
		Delete: resourceDocDBDBClusterParameterGroupDelete,
		CustomizeDiff: resourceDocDBDBClusterParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
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

func resourceDocDBDBClusterParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDocDBDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBClusterParameterGroupType, data, meta)
}

func resourceDocDBDBClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBClusterParameterGroupType, data, meta)
}
