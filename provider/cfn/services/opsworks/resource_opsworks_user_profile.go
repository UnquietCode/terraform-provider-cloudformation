// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
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

func ResourceOpsWorksUserProfile() *schema.Resource {
	return &schema.Resource{
		Read: resourceOpsWorksUserProfileRead,
		Create: resourceOpsWorksUserProfileCreate,
		Update: resourceOpsWorksUserProfileUpdate,
		Delete: resourceOpsWorksUserProfileDelete,
		CustomizeDiff: resourceOpsWorksUserProfileCustomizeDiff,
		
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
		},
	}
}

func resourceOpsWorksUserProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksUserProfileType, ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(opsWorksUserProfileType, ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(opsWorksUserProfileType, ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(opsWorksUserProfileType, data, meta)
}

func resourceOpsWorksUserProfileCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(opsWorksUserProfileType, data, meta)
}
