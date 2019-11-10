// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html

package amplify

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmplifyApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAmplifyAppCreate,
		Read:   resourceAmplifyAppRead,
		Update: resourceAmplifyAppUpdate,
		Delete: resourceAmplifyAppDelete,

		Schema: map[string]*schema.Schema{
			"auto_branch_creation_config": {
				Type: schema.TypeList,
				Elem: propertyAppAutoBranchCreationConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"oauth_token": {
				Type: schema.TypeString,
				Optional: true,
			},
			"repository": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyAppEnvironmentVariable(),
				Optional: true,
			},
			"access_token": {
				Type: schema.TypeString,
				Optional: true,
			},
			"build_spec": {
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_rules": {
				Type: schema.TypeList,
				Elem: propertyAppCustomRule(),
				Optional: true,
			},
			"basic_auth_config": {
				Type: schema.TypeList,
				Elem: propertyAppBasicAuthConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"iam_service_role": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAmplifyAppCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Amplify::App", data, meta)
}

func resourceAmplifyAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Amplify::App", data, meta)
}

func resourceAmplifyAppUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Amplify::App", data, meta)
}

func resourceAmplifyAppDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Amplify::App", data, meta)
}