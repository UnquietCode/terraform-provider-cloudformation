// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNeptuneDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNeptuneDBClusterParameterGroupCreate,
		Read:   resourceNeptuneDBClusterParameterGroupRead,
		Update: resourceNeptuneDBClusterParameterGroupUpdate,
		Delete: resourceNeptuneDBClusterParameterGroupDelete,

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
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceNeptuneDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBClusterParameterGroup", data, meta)
}

func resourceNeptuneDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBClusterParameterGroup", data, meta)
}

func resourceNeptuneDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBClusterParameterGroup", data, meta)
}

func resourceNeptuneDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBClusterParameterGroup", data, meta)
}