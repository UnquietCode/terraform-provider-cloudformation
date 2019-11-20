// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const codeCommitRepositoryType string = "AWS::CodeCommit::Repository"

func ResourceCodeCommitRepository() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodeCommitRepositoryRead,
		Create: resourceCodeCommitRepositoryCreate,
		Update: resourceCodeCommitRepositoryUpdate,
		Delete: resourceCodeCommitRepositoryDelete,
		CustomizeDiff: resourceCodeCommitRepositoryCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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
				Type: schema.TypeSet,
				Elem: propertyRepositoryCode(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCodeCommitRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeCommitRepositoryType, ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeCommitRepositoryType, ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeCommitRepositoryType, ResourceCodeCommitRepository(), data, meta)
}

func resourceCodeCommitRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeCommitRepositoryType, data, meta)
}

func resourceCodeCommitRepositoryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeCommitRepositoryType, data, meta)
}
