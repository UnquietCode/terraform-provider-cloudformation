// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html

package ecr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceECRRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceECRRepositoryCreate,
		Read:   resourceECRRepositoryRead,
		Update: resourceECRRepositoryUpdate,
		Delete: resourceECRRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"lifecycle_policy": {
				Type: schema.TypeList,
				Elem: propertyRepositoryLifecyclePolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"repository_policy_text": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceECRRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECR::Repository", ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECR::Repository", ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECR::Repository", ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECR::Repository", data, meta)
}