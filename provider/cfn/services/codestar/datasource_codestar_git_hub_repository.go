// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html

package codestar

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeStarGitHubRepositoryType string = "AWS::CodeStar::GitHubRepository"

func DatasourceCodeStarGitHubRepository() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodeStarGitHubRepositoryRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodeStarGitHubRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeStarGitHubRepositoryType, DatasourceCodeStarGitHubRepository(), data, meta)
}
