// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html

package ecr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eCRRepositoryType string = "AWS::ECR::Repository"

func DatasourceECRRepository() *schema.Resource {
	return &schema.Resource{
		Read: datasourceECRRepositoryRead,
		
		Schema: map[string]*schema.Schema{
			"lifecycle_policy": {
				Type: schema.TypeList,
				Elem: propertyRepositoryLifecyclePolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"repository_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"repository_policy_text": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"tags": misc.PropertyTags(),
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

func datasourceECRRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCRRepositoryType, DatasourceECRRepository(), data, meta)
}
