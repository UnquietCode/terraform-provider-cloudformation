// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-branch.html

package amplify

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amplifyBranchType string = "AWS::Amplify::Branch"

func ResourceAmplifyBranch() *schema.Resource {
	return &schema.Resource{
		Read: resourceAmplifyBranchRead,
		Create: resourceAmplifyBranchCreate,
		Update: resourceAmplifyBranchUpdate,
		Delete: resourceAmplifyBranchDelete,
		CustomizeDiff: resourceAmplifyBranchCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyBranchEnvironmentVariable(),
				Optional: true,
			},
			"app_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"pull_request_environment_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_pull_request_preview": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enable_auto_build": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"build_spec": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stage": {
				Type: schema.TypeString,
				Optional: true,
			},
			"branch_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"basic_auth_config": {
				Type: schema.TypeSet,
				Elem: propertyBranchBasicAuthConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAmplifyBranchRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amplifyBranchType, ResourceAmplifyBranch(), data, meta)
}

func resourceAmplifyBranchCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amplifyBranchType, ResourceAmplifyBranch(), data, meta)
}

func resourceAmplifyBranchUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amplifyBranchType, ResourceAmplifyBranch(), data, meta)
}

func resourceAmplifyBranchDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amplifyBranchType, data, meta)
}

func resourceAmplifyBranchCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amplifyBranchType, data, meta)
}
