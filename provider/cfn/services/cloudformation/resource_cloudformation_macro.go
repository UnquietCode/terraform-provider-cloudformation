// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFormationMacro() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFormationMacroCreate,
		Read:   resourceCloudFormationMacroRead,
		Update: resourceCloudFormationMacroUpdate,
		Delete: resourceCloudFormationMacroDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFormationMacroCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::Macro", ResourceCloudFormationMacro(), data, meta)
}

func resourceCloudFormationMacroRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::Macro", ResourceCloudFormationMacro(), data, meta)
}

func resourceCloudFormationMacroUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::Macro", ResourceCloudFormationMacro(), data, meta)
}

func resourceCloudFormationMacroDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::Macro", data, meta)
}