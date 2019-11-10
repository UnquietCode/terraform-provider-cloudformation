// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceParameter() *schema.Resource {
	return &schema.Resource{
		Create: resourceParameterCreate,
		Read:   resourceParameterRead,
		Update: resourceParameterUpdate,
		Delete: resourceParameterDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"policies": {
				Type: schema.TypeString,
				Required: false,
			},
			"allowed_pattern": {
				Type: schema.TypeString,
				Required: false,
			},
			"tier": {
				Type: schema.TypeString,
				Required: false,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
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

func resourceParameterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Parameter", data, meta)
}

func resourceParameterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Parameter", data, meta)
}

func resourceParameterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Parameter", data, meta)
}

func resourceParameterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Parameter", data, meta)
}