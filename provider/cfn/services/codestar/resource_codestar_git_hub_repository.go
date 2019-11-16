// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html

package codestar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeStarGitHubRepositoryType string = "AWS::CodeStar::GitHubRepository"

var codeStarGitHubRepositoryProperties map[string]string = map[string]string{
	"enable_issues": "EnableIssues",
	"repository_name": "RepositoryName",
	"repository_access_token": "RepositoryAccessToken",
	"repository_owner": "RepositoryOwner",
	"is_private": "IsPrivate",
	"code": "Code",
	"repository_description": "RepositoryDescription",
}

func ResourceCodeStarGitHubRepository() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodeStarGitHubRepositoryExists,
		Read: resourceCodeStarGitHubRepositoryRead,
		Create: resourceCodeStarGitHubRepositoryCreate,
		Update: resourceCodeStarGitHubRepositoryUpdate,
		Delete: resourceCodeStarGitHubRepositoryDelete,
		CustomizeDiff: resourceCodeStarGitHubRepositoryCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"enable_issues": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"repository_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"repository_access_token": {
				Type: schema.TypeString,
				Required: true,
			},
			"repository_owner": {
				Type: schema.TypeString,
				Required: true,
			},
			"is_private": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"code": {
				Type: schema.TypeList,
				Elem: propertyGitHubRepositoryCode(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCodeStarGitHubRepositoryExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodeStarGitHubRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeStarGitHubRepositoryType, ResourceCodeStarGitHubRepository(), data, meta)
}

func resourceCodeStarGitHubRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeStarGitHubRepositoryType, ResourceCodeStarGitHubRepository(), data, codeStarGitHubRepositoryProperties, meta)
}

func resourceCodeStarGitHubRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeStarGitHubRepositoryType, ResourceCodeStarGitHubRepository(), data, codeStarGitHubRepositoryProperties, meta)
}

func resourceCodeStarGitHubRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeStarGitHubRepositoryType, data, meta)
}

func resourceCodeStarGitHubRepositoryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeStarGitHubRepositoryType, data, meta)
}
