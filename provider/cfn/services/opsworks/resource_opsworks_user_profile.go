// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksUserProfile() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksUserProfileExists,
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
			},
		},
	}
}

func resourceOpsWorksUserProfileExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksUserProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::UserProfile", ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::UserProfile", ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::UserProfile", ResourceOpsWorksUserProfile(), data, meta)
}

func resourceOpsWorksUserProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::UserProfile", data, meta)
}

func resourceOpsWorksUserProfileCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::OpsWorks::UserProfile", data, meta)
}

