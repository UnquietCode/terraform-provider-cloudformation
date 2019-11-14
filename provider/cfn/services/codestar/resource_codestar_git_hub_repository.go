// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html

package codestar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
	return plugin.ResourceRead("AWS::CodeStar::GitHubRepository", ResourceCodeStarGitHubRepository(), data, meta)
}

func resourceCodeStarGitHubRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeStar::GitHubRepository", ResourceCodeStarGitHubRepository(), data, meta)
}

func resourceCodeStarGitHubRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeStar::GitHubRepository", ResourceCodeStarGitHubRepository(), data, meta)
}

func resourceCodeStarGitHubRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeStar::GitHubRepository", data, meta)
}

func resourceCodeStarGitHubRepositoryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
