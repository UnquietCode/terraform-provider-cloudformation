// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-coderepository.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSageMakerCodeRepository() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSageMakerCodeRepositoryExists,
		Read:   resourceSageMakerCodeRepositoryRead,
		Create: resourceSageMakerCodeRepositoryCreate,
		Update: resourceSageMakerCodeRepositoryUpdate,
		Delete: resourceSageMakerCodeRepositoryDelete,
		CustomizeDiff: resourceSageMakerCodeRepositoryCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"code_repository_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"git_config": {
				Type: schema.TypeList,
				Elem: propertyCodeRepositoryGitConfig(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSageMakerCodeRepositoryExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerCodeRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::CodeRepository", ResourceSageMakerCodeRepository(), data, meta)
}

func resourceSageMakerCodeRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::CodeRepository", ResourceSageMakerCodeRepository(), data, meta)
}

func resourceSageMakerCodeRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::CodeRepository", ResourceSageMakerCodeRepository(), data, meta)
}

func resourceSageMakerCodeRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::CodeRepository", data, meta)
}

func resourceSageMakerCodeRepositoryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}