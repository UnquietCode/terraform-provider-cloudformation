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

func ResourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"auto_branch_creation_config": {
				Type: schema.TypeList,
				Elem: propertyAppAutoBranchCreationConfig(),
				Required: false,
				MaxItems: 1,
			},
			"oauth_token": {
				Type: schema.TypeString,
				Required: false,
			},
			"repository": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Required: false,
			},
			"access_token": {
				Type: schema.TypeString,
				Required: false,
			},
			"build_spec": {
				Type: schema.TypeString,
				Required: false,
			},
			"custom_rules": {
				Type: schema.TypeList,
				Elem: propertyAppCustomRule(),
				Required: false,
			},
			"basic_auth_config": {
				Type: schema.TypeList,
				Elem: propertyAppBasicAuthConfig(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"iam_service_role": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceAppCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Amplify::App", data, meta)
}

func resourceAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Amplify::App", data, meta)
}

func resourceAppUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Amplify::App", data, meta)
}

func resourceAppDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Amplify::App", data, meta)
}