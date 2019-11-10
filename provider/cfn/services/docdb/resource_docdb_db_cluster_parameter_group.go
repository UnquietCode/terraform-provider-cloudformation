// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceDocDBDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocDBDBClusterParameterGroupCreate,
		Read:   resourceDocDBDBClusterParameterGroupRead,
		Update: resourceDocDBDBClusterParameterGroupUpdate,
		Delete: resourceDocDBDBClusterParameterGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDocDBDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DocDB::DBClusterParameterGroup", ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DocDB::DBClusterParameterGroup", ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DocDB::DBClusterParameterGroup", ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DocDB::DBClusterParameterGroup", data, meta)
}