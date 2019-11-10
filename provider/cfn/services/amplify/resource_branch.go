// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package amplify

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBranch() *schema.Resource {
	return &schema.Resource{
		Create: resourceBranchCreate,
		Read:   resourceBranchRead,
		Update: resourceBranchUpdate,
		Delete: resourceBranchDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyBranchEnvironmentVariable(),
				Required: false,
			},
			"app_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pull_request_environment_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"enable_pull_request_preview": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enable_auto_build": {
				Type: schema.TypeBool,
				Required: false,
			},
			"build_spec": {
				Type: schema.TypeString,
				Required: false,
			},
			"stage": {
				Type: schema.TypeString,
				Required: false,
			},
			"branch_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"basic_auth_config": {
				Type: schema.TypeList,
				Elem: propertyBranchBasicAuthConfig(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceBranchCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Amplify::Branch", data, meta)
}

func resourceBranchRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Amplify::Branch", data, meta)
}

func resourceBranchUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Amplify::Branch", data, meta)
}

func resourceBranchDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Amplify::Branch", data, meta)
}