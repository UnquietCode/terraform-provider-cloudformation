// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codestar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGitHubRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceGitHubRepositoryCreate,
		Read:   resourceGitHubRepositoryRead,
		Update: resourceGitHubRepositoryUpdate,
		Delete: resourceGitHubRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"enable_issues": {
				Type: schema.TypeBool,
				Required: false,
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
				Required: false,
			},
			"code": {
				Type: schema.TypeList,
				Elem: propertyGitHubRepositoryCode(),
				Required: false,
				MaxItems: 1,
			},
			"repository_description": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceGitHubRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeStar::GitHubRepository", data, meta)
}

func resourceGitHubRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeStar::GitHubRepository", data, meta)
}

func resourceGitHubRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeStar::GitHubRepository", data, meta)
}

func resourceGitHubRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeStar::GitHubRepository", data, meta)
}