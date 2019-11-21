// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html

package amplify

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amplifyAppType string = "AWS::Amplify::App"

func ResourceAmplifyApp() *schema.Resource {
	return &schema.Resource{
		Read: resourceAmplifyAppRead,
		Create: resourceAmplifyAppCreate,
		Update: resourceAmplifyAppUpdate,
		Delete: resourceAmplifyAppDelete,
		CustomizeDiff: resourceAmplifyAppCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_branch_creation_config": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyAppBasicAuthConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"iam_service_role": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAmplifyAppRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amplifyAppType, ResourceAmplifyApp(), data, meta)
}

func resourceAmplifyAppCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amplifyAppType, ResourceAmplifyApp(), data, meta)
}

func resourceAmplifyAppUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amplifyAppType, ResourceAmplifyApp(), data, meta)
}

func resourceAmplifyAppDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amplifyAppType, data, meta)
}

func resourceAmplifyAppCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amplifyAppType, data, meta)
}
