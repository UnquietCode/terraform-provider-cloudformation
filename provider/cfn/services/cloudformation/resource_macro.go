// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMacro() *schema.Resource {
	return &schema.Resource{
		Create: resourceMacroCreate,
		Read:   resourceMacroRead,
		Update: resourceMacroUpdate,
		Delete: resourceMacroDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"log_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMacroCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::Macro", data, meta)
}

func resourceMacroRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::Macro", data, meta)
}

func resourceMacroUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::Macro", data, meta)
}

func resourceMacroDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::Macro", data, meta)
}