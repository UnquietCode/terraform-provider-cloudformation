// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBClusterParameterGroupCreate,
		Read:   resourceDBClusterParameterGroupRead,
		Update: resourceDBClusterParameterGroupUpdate,
		Delete: resourceDBClusterParameterGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBClusterParameterGroup", data, meta)
}

func resourceDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBClusterParameterGroup", data, meta)
}

func resourceDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBClusterParameterGroup", data, meta)
}

func resourceDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBClusterParameterGroup", data, meta)
}