// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dax

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDAXParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDAXParameterGroupCreate,
		Read:   resourceDAXParameterGroupRead,
		Update: resourceDAXParameterGroupUpdate,
		Delete: resourceDAXParameterGroupDelete,

		Schema: map[string]*schema.Schema{
			"parameter_name_values": {
				Type: schema.TypeMap,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceDAXParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DAX::ParameterGroup", data, meta)
}

func resourceDAXParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DAX::ParameterGroup", data, meta)
}

func resourceDAXParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DAX::ParameterGroup", data, meta)
}

func resourceDAXParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DAX::ParameterGroup", data, meta)
}