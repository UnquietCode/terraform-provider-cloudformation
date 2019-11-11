// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html

package codecommit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeCommitRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeCommitRepositoryCreate,
		Read:   resourceCodeCommitRepositoryRead,
		Update: resourceCodeCommitRepositoryUpdate,
		Delete: resourceCodeCommitRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"clone_url_http": {
				Type: schema.TypeString,
				Computed: true,
			},
			"clone_url_ssh": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"repository_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"triggers": {
				Type: schema.TypeList,
				Elem: propertyRepositoryRepositoryTrigger(),
				Optional: true,
			},
			"code": {
				Type: schema.TypeList,
				Elem: propertyRepositoryCode(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceCodeCommitRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeCommit::Repository", ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeCommit::Repository", ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeCommit::Repository", ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeCommit::Repository", data, meta)
}