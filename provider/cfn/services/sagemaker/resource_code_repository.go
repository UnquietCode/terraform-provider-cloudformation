// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeRepositoryCreate,
		Read:   resourceCodeRepositoryRead,
		Update: resourceCodeRepositoryUpdate,
		Delete: resourceCodeRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"code_repository_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"git_config": {
				Type: schema.TypeList,
				Elem: propertyCodeRepositoryGitConfig(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceCodeRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::CodeRepository", data, meta)
}

func resourceCodeRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::CodeRepository", data, meta)
}

func resourceCodeRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::CodeRepository", data, meta)
}

func resourceCodeRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::CodeRepository", data, meta)
}