// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html

package opsworks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksUserProfileType string = "AWS::OpsWorks::UserProfile"

func DatasourceOpsWorksUserProfile() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOpsWorksUserProfileRead,
		
		Schema: map[string]*schema.Schema{
			"allow_self_management": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"iam_user_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"ssh_public_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssh_username": {
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

func datasourceOpsWorksUserProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksUserProfileType, DatasourceOpsWorksUserProfile(), data, meta)
}
