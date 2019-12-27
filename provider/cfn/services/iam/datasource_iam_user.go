// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html

package iam

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMUserType string = "AWS::IAM::User"

func DatasourceIAMUser() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIAMUserRead,
		
		Schema: map[string]*schema.Schema{
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"login_profile": {
				Type: schema.TypeList,
				Elem: propertyUserLoginProfile(),
				Optional: true,
				MaxItems: 1,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"permissions_boundary": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyUserPolicy(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"user_name": {
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

func datasourceIAMUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMUserType, DatasourceIAMUser(), data, meta)
}
