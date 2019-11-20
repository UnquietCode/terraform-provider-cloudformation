// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const eCRRepositoryType string = "AWS::ECR::Repository"

func ResourceECRRepository() *schema.Resource {
	return &schema.Resource{
		Read: resourceECRRepositoryRead,
		Create: resourceECRRepositoryCreate,
		Update: resourceECRRepositoryUpdate,
		Delete: resourceECRRepositoryDelete,
		CustomizeDiff: resourceECRRepositoryCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"lifecycle_policy": {
				Type: schema.TypeSet,
				Elem: propertyRepositoryLifecyclePolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"repository_policy_text": {
				Type: schema.TypeMap,
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

func resourceECRRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCRRepositoryType, ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eCRRepositoryType, ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eCRRepositoryType, ResourceECRRepository(), data, meta)
}

func resourceECRRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eCRRepositoryType, data, meta)
}

func resourceECRRepositoryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eCRRepositoryType, data, meta)
}
