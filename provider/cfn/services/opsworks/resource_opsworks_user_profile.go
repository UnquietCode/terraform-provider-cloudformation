// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksUserProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpsWorksUserProfileCreate,
		Read:   resourceOpsWorksUserProfileRead,
		Update: resourceOpsWorksUserProfileUpdate,
		Delete: resourceOpsWorksUserProfileDelete,

		Schema: map[string]*schema.Schema{
			"allow_self_management": {
				Type: schema.TypeBool,
				Required: false,
			},
			"iam_user_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"ssh_username": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceOpsWorksUserProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::UserProfile", data, meta)
}

func resourceOpsWorksUserProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::UserProfile", data, meta)
}

func resourceOpsWorksUserProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::UserProfile", data, meta)
}

func resourceOpsWorksUserProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::UserProfile", data, meta)
}